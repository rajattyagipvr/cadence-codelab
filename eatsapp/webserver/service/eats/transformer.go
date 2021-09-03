package eats

import (
	"errors"
	"path"
	"strconv"
	"context"
	"go.uber.org/cadence/client"
	s "go.uber.org/cadence/.gen/go/shared"
)

type (
	// TransformFunc type defining the signature of transform function.
	transformFunc func(event *s.HistoryEvent, tasks *TaskGroup) error

	// TaskGroupExecution implements object to transform a workflow history into a TaskGroup.
	TaskGroupExecution struct {
		client       client.Client
		transformers map[s.EventType]transformFunc
	}
)

// NewTaskGroupExecution returns a new instanc of TaskGroupExecution.
func NewTaskGroupExecution(c client.Client) *TaskGroupExecution {
	obj := &TaskGroupExecution{
		client:       c,
		transformers: make(map[s.EventType]transformFunc),
	}

	obj.transformers[s.EventTypeWorkflowExecutionStarted] = obj.tfWorkflowExecutionStarted
	obj.transformers[s.EventTypeWorkflowExecutionCompleted] = obj.tfWorkflowExecutionCompleted
	obj.transformers[s.EventTypeWorkflowExecutionFailed] = obj.tfWorkflowExecutionFailed

	obj.transformers[s.EventTypeActivityTaskScheduled] = obj.tfActivityTaskScheduled
	obj.transformers[s.EventTypeActivityTaskStarted] = obj.tfActivityTaskStarted
	obj.transformers[s.EventTypeActivityTaskCompleted] = obj.tfActivityTaskCompleted
	obj.transformers[s.EventTypeActivityTaskFailed] = obj.tfActivityTaskFailed
	obj.transformers[s.EventTypeActivityTaskTimedOut] = obj.tfActivityTaskTimedOut

	obj.transformers[s.EventTypeStartChildWorkflowExecutionInitiated] = obj.tfStartChildWorkflowExecutionInitiated
	obj.transformers[s.EventTypeChildWorkflowExecutionStarted] = obj.tfChildWorkflowExecutionStarted
	obj.transformers[s.EventTypeChildWorkflowExecutionCompleted] = obj.tfChildWorkflowExecutionCompleted
	obj.transformers[s.EventTypeChildWorkflowExecutionFailed] = obj.tfChildWorkflowExecutionFailed
	obj.transformers[s.EventTypeChildWorkflowExecutionTimedOut] = obj.tfChildWorkflowExecutionTimedOut

	obj.transformers[s.EventTypeTimerStarted] = obj.tfTimerStarted
	obj.transformers[s.EventTypeTimerFired] = obj.tfTimerFired
	obj.transformers[s.EventTypeTimerCanceled] = obj.tfTimerCanceled
	return obj
}

// Transform converts a workflow execution history into a TaskGroup structure.
func (h *TaskGroupExecution) Transform(workflowID string, runID string) (*TaskGroup, error) {
	tasks := &TaskGroup{
		ID:      workflowID,
		RunID:   runID,
		Tasks:   make([]*Task, 0),
		TaskMap: make(map[int64]*Task),
	}
	ctx := context.Background()
	history, err := h.client.GetWorkflowHistory(ctx, workflowID, runID, false, s.HistoryEventFilterTypeAllEvent), error(nil)
	if err != nil {
		return nil, err
	}
	//tasks.History = history
	//rajat
	
	for _ = history; history.HasNext(); {
		value, _ := history.Next()
		tasks.History.Events = append(tasks.History.Events, value) 
	}

	
	for _, event := range tasks.History.Events {
		transFunc, found := h.transformers[*event.EventType]
		if !found {
			continue
		}

		err := transFunc(event, tasks)
		if err != nil {
			return nil, err
		}
	}

	tasks.TaskMap = nil
	return tasks, nil
}

func (h *TaskGroupExecution) tfActivityTaskScheduled(event *s.HistoryEvent, tasks *TaskGroup) error {
	name := event.ActivityTaskScheduledEventAttributes.ActivityType.Name
	return h.createTask(event, name, tasks)
}

func (h *TaskGroupExecution) tfActivityTaskStarted(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ActivityTaskStartedEventAttributes.ScheduledEventId
	return h.setTaskStatus(tasks, id, "r")
}

func (h *TaskGroupExecution) tfActivityTaskCompleted(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ActivityTaskCompletedEventAttributes.ScheduledEventId
	return h.setTaskStatus(tasks, id, "c")
}

func (h *TaskGroupExecution) tfActivityTaskFailed(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ActivityTaskFailedEventAttributes.ScheduledEventId
	return h.setTaskStatus(tasks, id, "f")
}

func (h *TaskGroupExecution) tfActivityTaskTimedOut(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ActivityTaskTimedOutEventAttributes.ScheduledEventId
	return h.setTaskStatus(tasks, id, "t")
}

func (h *TaskGroupExecution) tfStartChildWorkflowExecutionInitiated(event *s.HistoryEvent, tasks *TaskGroup) error {
	name := event.StartChildWorkflowExecutionInitiatedEventAttributes.WorkflowType.Name
	return h.createTask(event, name, tasks)
}

func (h *TaskGroupExecution) tfChildWorkflowExecutionStarted(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ChildWorkflowExecutionStartedEventAttributes.InitiatedEventId
	task, found := tasks.TaskMap[id]
	if !found {
		return errors.New("Could not find ActivityTaskScheduled event: " + strconv.FormatInt(id, 10))
	}

	execution := event.ChildWorkflowExecutionStartedEventAttributes.WorkflowExecution
	taskGroup, err := h.Transform(*execution.WorkflowId, *execution.RunId)
	if err != nil {
		return err
	}

	task.SubTasks = taskGroup.Tasks
	task.Status = "r"
	return nil
}

func (h *TaskGroupExecution) tfChildWorkflowExecutionCompleted(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ChildWorkflowExecutionCompletedEventAttributes.InitiatedEventId
	return h.setTaskStatus(tasks, id, "c")
}

func (h *TaskGroupExecution) tfChildWorkflowExecutionFailed(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ChildWorkflowExecutionFailedEventAttributes.InitiatedEventId
	return h.setTaskStatus(tasks, id, "f")
}

func (h *TaskGroupExecution) tfChildWorkflowExecutionTimedOut(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.ChildWorkflowExecutionTimedOutEventAttributes.InitiatedEventId
	return h.setTaskStatus(tasks, id, "t")
}

func (h *TaskGroupExecution) tfTimerStarted(event *s.HistoryEvent, tasks *TaskGroup) error {
	name := "timer.WaitForDeadline"
	h.createTask(event, &name, tasks)
	return h.setTaskStatus(tasks, *event.EventId, "r")
}

func (h *TaskGroupExecution) tfTimerCanceled(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.TimerCanceledEventAttributes.StartedEventId
	return h.setTaskStatus(tasks, id, "ca")
}

func (h *TaskGroupExecution) tfTimerFired(event *s.HistoryEvent, tasks *TaskGroup) error {
	id := *event.TimerFiredEventAttributes.StartedEventId
	return h.setTaskStatus(tasks, id, "c")
}

func (h *TaskGroupExecution) tfWorkflowExecutionStarted(event *s.HistoryEvent, tasks *TaskGroup) error {
	tasks.Status = "r"
	return nil
}

func (h *TaskGroupExecution) tfWorkflowExecutionCompleted(event *s.HistoryEvent, tasks *TaskGroup) error {
	tasks.Status = "c"
	return nil
}

func (h *TaskGroupExecution) tfWorkflowExecutionFailed(event *s.HistoryEvent, tasks *TaskGroup) error {
	tasks.Status = "f"
	return nil
}

func (h *TaskGroupExecution) createTask(event *s.HistoryEvent, name *string, tasks *TaskGroup) error {
	task := &Task{
		ID:     *event.EventId,
		Name:   path.Ext(*name)[1:],
		Status: "s",
	}

	tasks.TaskMap[task.ID] = task
	tasks.Tasks = append(tasks.Tasks, task)

	return nil
}

func (h *TaskGroupExecution) setTaskStatus(tasks *TaskGroup, id int64, status TaskStatus) error {
	task, found := tasks.TaskMap[id]
	if !found {
		return errors.New("Could not find ActivityTaskScheduled event: " + strconv.FormatInt(id, 10))
	}

	task.Status = status
	return nil
}
