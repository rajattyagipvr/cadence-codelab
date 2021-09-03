// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

namespace java com.uber.cadence

exception BadRequestError {
  1: required string message
}

exception InternalServiceError {
  1: required string message
}

exception DomainAlreadyExistsError {
  1: required string message
}

exception WorkflowExecutionAlreadyStartedError {
  10: optional string message
  20: optional string startRequestId
  30: optional string runId
}

exception EntityNotExistsError {
  1: required string message
}

exception ServiceBusyError {
  1: required string message
}

enum DomainStatus {
  REGISTERED,
  DEPRECATED,
  DELETED,
}

enum TimeoutType {
  START_TO_CLOSE,
  SCHEDULE_TO_START,
  SCHEDULE_TO_CLOSE,
  HEARTBEAT,
}

enum DecisionType {
  ScheduleActivityTask,
  RequestCancelActivityTask,
  StartTimer,
  CompleteWorkflowExecution,
  FailWorkflowExecution,
  CancelTimer,
  CancelWorkflowExecution,
  RequestCancelExternalWorkflowExecution,
  RecordMarker,
  ContinueAsNewWorkflowExecution,
  StartChildWorkflowExecution,
}

enum EventType {
  WorkflowExecutionStarted,
  WorkflowExecutionCompleted,
  WorkflowExecutionFailed,
  WorkflowExecutionTimedOut,
  DecisionTaskScheduled,
  DecisionTaskStarted,
  DecisionTaskCompleted,
  DecisionTaskTimedOut
  DecisionTaskFailed,
  ActivityTaskScheduled,
  ActivityTaskStarted,
  ActivityTaskCompleted,
  ActivityTaskFailed,
  ActivityTaskTimedOut,
  ActivityTaskCancelRequested,
  RequestCancelActivityTaskFailed,
  ActivityTaskCanceled,
  TimerStarted,
  TimerFired,
  CancelTimerFailed,
  TimerCanceled,
  WorkflowExecutionCancelRequested,
  WorkflowExecutionCanceled,
  RequestCancelExternalWorkflowExecutionInitiated,
  RequestCancelExternalWorkflowExecutionFailed,
  ExternalWorkflowExecutionCancelRequested,
  MarkerRecorded,
  WorkflowExecutionSignaled,
  WorkflowExecutionTerminated,
  WorkflowExecutionContinuedAsNew,
  StartChildWorkflowExecutionInitiated,
  StartChildWorkflowExecutionFailed,
  ChildWorkflowExecutionStarted,
  ChildWorkflowExecutionCompleted,
  ChildWorkflowExecutionFailed,
  ChildWorkflowExecutionCanceled,
  ChildWorkflowExecutionTimedOut,
  ChildWorkflowExecutionTerminated,
}

enum DecisionTaskFailedCause {
  UNHANDLED_DECISION,
  BAD_SCHEDULE_ACTIVITY_ATTRIBUTES,
  BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES,
  BAD_START_TIMER_ATTRIBUTES,
  BAD_CANCEL_TIMER_ATTRIBUTES,
  BAD_RECORD_MARKER_ATTRIBUTES,
  BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES,
  BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES,
  BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES,
  BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES,
  BAD_CONTINUE_AS_NEW_ATTRIBUTES,
}

enum CancelExternalWorkflowExecutionFailedCause {
  UNKNOWN_EXTERNAL_WORKFLOW_EXECUTION,
}

enum ChildWorkflowExecutionFailedCause {
  WORKFLOW_ALREADY_RUNNING,
}

enum WorkflowExecutionCloseStatus {
  COMPLETED,
  FAILED,
  CANCELED,
  TERMINATED,
  CONTINUED_AS_NEW,
  TIMED_OUT,
}

enum ChildPolicy {
  TERMINATE,
  REQUEST_CANCEL,
  ABANDON,
}

struct WorkflowType {
  10: optional string name
}

struct ActivityType {
  10: optional string name
}

struct TaskList {
  10: optional string name
}

struct WorkflowExecution {
  10: optional string workflowId
  20: optional string runId
}

struct WorkflowExecutionInfo {
  10: optional WorkflowExecution execution
  20: optional WorkflowType type
  30: optional i64 (js.type = "Long") startTime
  40: optional i64 (js.type = "Long") closeTime
  50: optional WorkflowExecutionCloseStatus closeStatus
  60: optional i64 (js.type = "Long") historyLength
}

struct ScheduleActivityTaskDecisionAttributes {
  10: optional string activityId
  20: optional ActivityType activityType
  25: optional string domain
  30: optional TaskList taskList
  40: optional binary input
  45: optional i32 scheduleToCloseTimeoutSeconds
  50: optional i32 scheduleToStartTimeoutSeconds
  55: optional i32 startToCloseTimeoutSeconds
  60: optional i32 heartbeatTimeoutSeconds
}

struct RequestCancelActivityTaskDecisionAttributes {
  10: optional string activityId
}

struct StartTimerDecisionAttributes {
  10: optional string timerId
  20: optional i64 (js.type = "Long") startToFireTimeoutSeconds
}

struct CompleteWorkflowExecutionDecisionAttributes {
  10: optional binary result
}

struct FailWorkflowExecutionDecisionAttributes {
  10: optional string reason
  20: optional binary details
}

struct CancelTimerDecisionAttributes {
  10: optional string timerId
}

struct CancelWorkflowExecutionDecisionAttributes {
  10: optional binary details
}

struct RequestCancelExternalWorkflowExecutionDecisionAttributes {
  10: optional string domain
  20: optional string workflowId
  30: optional string runId
  40: optional binary control
}

struct RecordMarkerDecisionAttributes {
  10: optional string markerName
  20: optional binary details
}

struct ContinueAsNewWorkflowExecutionDecisionAttributes {
  10: optional WorkflowType workflowType
  20: optional TaskList taskList
  30: optional binary input
  40: optional i32 executionStartToCloseTimeoutSeconds
  50: optional i32 taskStartToCloseTimeoutSeconds
}

struct StartChildWorkflowExecutionDecisionAttributes {
  10: optional string domain
  20: optional string workflowId
  30: optional WorkflowType workflowType
  40: optional TaskList taskList
  50: optional binary input
  60: optional i32 executionStartToCloseTimeoutSeconds
  70: optional i32 taskStartToCloseTimeoutSeconds
  80: optional ChildPolicy childPolicy
  90: optional binary control
}

struct Decision {
  10:  optional DecisionType decisionType
  20:  optional ScheduleActivityTaskDecisionAttributes scheduleActivityTaskDecisionAttributes
  25:  optional StartTimerDecisionAttributes startTimerDecisionAttributes
  30:  optional CompleteWorkflowExecutionDecisionAttributes completeWorkflowExecutionDecisionAttributes
  35:  optional FailWorkflowExecutionDecisionAttributes failWorkflowExecutionDecisionAttributes
  40:  optional RequestCancelActivityTaskDecisionAttributes requestCancelActivityTaskDecisionAttributes
  50:  optional CancelTimerDecisionAttributes cancelTimerDecisionAttributes
  60:  optional CancelWorkflowExecutionDecisionAttributes cancelWorkflowExecutionDecisionAttributes
  70:  optional RequestCancelExternalWorkflowExecutionDecisionAttributes requestCancelExternalWorkflowExecutionDecisionAttributes
  80:  optional RecordMarkerDecisionAttributes recordMarkerDecisionAttributes
  90:  optional ContinueAsNewWorkflowExecutionDecisionAttributes continueAsNewWorkflowExecutionDecisionAttributes
  100: optional StartChildWorkflowExecutionDecisionAttributes startChildWorkflowExecutionDecisionAttributes
}

struct WorkflowExecutionStartedEventAttributes {
  10: optional WorkflowType workflowType
  20: optional TaskList taskList
  30: optional binary input
  40: optional i32 executionStartToCloseTimeoutSeconds
  50: optional i32 taskStartToCloseTimeoutSeconds
  60: optional string identity
}

struct WorkflowExecutionCompletedEventAttributes {
  10: optional binary result
  20: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct WorkflowExecutionFailedEventAttributes {
  10: optional string reason
  20: optional binary details
  30: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct WorkflowExecutionTimedOutEventAttributes {
  10: optional TimeoutType timeoutType
}

struct WorkflowExecutionContinuedAsNewEventAttributes {
  10: optional string newExecutionRunId
  20: optional WorkflowType workflowType
  30: optional TaskList taskList
  40: optional binary input
  50: optional i32 executionStartToCloseTimeoutSeconds
  60: optional i32 taskStartToCloseTimeoutSeconds
  70: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct DecisionTaskScheduledEventAttributes {
  10: optional TaskList taskList
  20: optional i32 startToCloseTimeoutSeconds
}

struct DecisionTaskStartedEventAttributes {
  10: optional i64 (js.type = "Long") scheduledEventId
  20: optional string identity
  30: optional string requestId
}

struct DecisionTaskCompletedEventAttributes {
  10: optional binary executionContext
  20: optional i64 (js.type = "Long") scheduledEventId
  30: optional i64 (js.type = "Long") startedEventId
  40: optional string identity
}

struct DecisionTaskTimedOutEventAttributes {
  10: optional i64 (js.type = "Long") scheduledEventId
  20: optional i64 (js.type = "Long") startedEventId
  30: optional TimeoutType timeoutType
}

struct DecisionTaskFailedEventAttributes {
  10: optional i64 (js.type = "Long") scheduledEventId
  20: optional i64 (js.type = "Long") startedEventId
  30: optional DecisionTaskFailedCause cause
  40: optional string identity
}

struct ActivityTaskScheduledEventAttributes {
  10: optional string activityId
  20: optional ActivityType activityType
  25: optional string domain
  30: optional TaskList taskList
  40: optional binary input
  45: optional i32 scheduleToCloseTimeoutSeconds
  50: optional i32 scheduleToStartTimeoutSeconds
  55: optional i32 startToCloseTimeoutSeconds
  60: optional i32 heartbeatTimeoutSeconds
  90: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct ActivityTaskStartedEventAttributes {
  10: optional i64 (js.type = "Long") scheduledEventId
  20: optional string identity
  30: optional string requestId
}

struct ActivityTaskCompletedEventAttributes {
  10: optional binary result
  20: optional i64 (js.type = "Long") scheduledEventId
  30: optional i64 (js.type = "Long") startedEventId
  40: optional string identity
}

struct ActivityTaskFailedEventAttributes {
  10: optional string reason
  20: optional binary details
  30: optional i64 (js.type = "Long") scheduledEventId
  40: optional i64 (js.type = "Long") startedEventId
  50: optional string identity
}

struct ActivityTaskTimedOutEventAttributes {
  05: optional binary details
  10: optional i64 (js.type = "Long") scheduledEventId
  20: optional i64 (js.type = "Long") startedEventId
  30: optional TimeoutType timeoutType
}

struct ActivityTaskCancelRequestedEventAttributes {
  10: optional string activityId
  20: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct RequestCancelActivityTaskFailedEventAttributes{
  10: optional string activityId
  20: optional string cause
  30: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct ActivityTaskCanceledEventAttributes {
  10: optional binary details
  20: optional i64 (js.type = "Long") latestCancelRequestedEventId
  30: optional i64 (js.type = "Long") scheduledEventId
  40: optional i64 (js.type = "Long") startedEventId
  50: optional string identity
}

struct TimerStartedEventAttributes {
  10: optional string timerId
  20: optional i64 (js.type = "Long") startToFireTimeoutSeconds
  30: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct TimerFiredEventAttributes {
  10: optional string timerId
  20: optional i64 (js.type = "Long") startedEventId
}

struct TimerCanceledEventAttributes {
  10: optional string timerId
  20: optional i64 (js.type = "Long") startedEventId
  30: optional i64 (js.type = "Long") decisionTaskCompletedEventId
  40: optional string identity
}

struct CancelTimerFailedEventAttributes {
  10: optional string timerId
  20: optional string cause
  30: optional i64 (js.type = "Long") decisionTaskCompletedEventId
  40: optional string identity
}

struct WorkflowExecutionCancelRequestedEventAttributes {
  10: optional string cause
  20: optional i64 (js.type = "Long") externalInitiatedEventId
  30: optional WorkflowExecution externalWorkflowExecution
  40: optional string identity
}

struct WorkflowExecutionCanceledEventAttributes {
  10: optional i64 (js.type = "Long") decisionTaskCompletedEventId
  20: optional binary details
}

struct MarkerRecordedEventAttributes {
  10: optional string markerName
  20: optional binary details
  30: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct WorkflowExecutionSignaledEventAttributes {
  10: optional string signalName
  20: optional binary input
  30: optional string identity
}

struct WorkflowExecutionTerminatedEventAttributes {
  10: optional string reason
  20: optional binary details
  30: optional string identity
}

struct RequestCancelExternalWorkflowExecutionInitiatedEventAttributes {
  10: optional i64 (js.type = "Long") decisionTaskCompletedEventId
  20: optional string domain
  30: optional WorkflowExecution workflowExecution
  40: optional binary control
}

struct RequestCancelExternalWorkflowExecutionFailedEventAttributes {
  10: optional CancelExternalWorkflowExecutionFailedCause cause
  20: optional i64 (js.type = "Long") decisionTaskCompletedEventId
  30: optional string domain
  40: optional WorkflowExecution workflowExecution
  50: optional i64 (js.type = "Long") initiatedEventId
  60: optional binary control
}

struct ExternalWorkflowExecutionCancelRequestedEventAttributes {
  10: optional i64 (js.type = "Long") initiatedEventId
  20: optional string domain
  30: optional WorkflowExecution workflowExecution
}

struct StartChildWorkflowExecutionInitiatedEventAttributes {
  10:  optional string domain
  20:  optional string workflowId
  30:  optional WorkflowType workflowType
  40:  optional TaskList taskList
  50:  optional binary input
  60:  optional i32 executionStartToCloseTimeoutSeconds
  70:  optional i32 taskStartToCloseTimeoutSeconds
  80:  optional ChildPolicy childPolicy
  90:  optional binary control
  100: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct StartChildWorkflowExecutionFailedEventAttributes {
  10: optional string domain
  20: optional string workflowId
  30: optional WorkflowType workflowType
  40: optional ChildWorkflowExecutionFailedCause cause
  50: optional binary control
  60: optional i64 (js.type = "Long") initiatedEventId
  70: optional i64 (js.type = "Long") decisionTaskCompletedEventId
}

struct ChildWorkflowExecutionStartedEventAttributes {
  10: optional string domain
  20: optional i64 (js.type = "Long") initiatedEventId
  30: optional WorkflowExecution workflowExecution
  40: optional WorkflowType workflowType
}

struct ChildWorkflowExecutionCompletedEventAttributes {
  10: optional binary result
  20: optional string domain
  30: optional WorkflowExecution workflowExecution
  40: optional WorkflowType workflowType
  50: optional i64 (js.type = "Long") initiatedEventId
  60: optional i64 (js.type = "Long") startedEventId
}

struct ChildWorkflowExecutionFailedEventAttributes {
  10: optional string reason
  20: optional binary details
  30: optional string domain
  40: optional WorkflowExecution workflowExecution
  50: optional WorkflowType workflowType
  60: optional i64 (js.type = "Long") initiatedEventId
  70: optional i64 (js.type = "Long") startedEventId
}

struct ChildWorkflowExecutionCanceledEventAttributes {
  10: optional binary details
  20: optional string domain
  30: optional WorkflowExecution workflowExecution
  40: optional WorkflowType workflowType
  50: optional i64 (js.type = "Long") initiatedEventId
  60: optional i64 (js.type = "Long") startedEventId
}

struct ChildWorkflowExecutionTimedOutEventAttributes {
  10: optional TimeoutType timeoutType
  20: optional string domain
  30: optional WorkflowExecution workflowExecution
  40: optional WorkflowType workflowType
  50: optional i64 (js.type = "Long") initiatedEventId
  60: optional i64 (js.type = "Long") startedEventId
}

struct ChildWorkflowExecutionTerminatedEventAttributes {
  10: optional string domain
  20: optional WorkflowExecution workflowExecution
  30: optional WorkflowType workflowType
  40: optional i64 (js.type = "Long") initiatedEventId
  50: optional i64 (js.type = "Long") startedEventId
}

struct HistoryEvent {
  10:  optional i64 (js.type = "Long") eventId
  20:  optional i64 (js.type = "Long") timestamp
  30:  optional EventType eventType
  40:  optional WorkflowExecutionStartedEventAttributes workflowExecutionStartedEventAttributes
  50:  optional WorkflowExecutionCompletedEventAttributes workflowExecutionCompletedEventAttributes
  60:  optional WorkflowExecutionFailedEventAttributes workflowExecutionFailedEventAttributes
  70:  optional WorkflowExecutionTimedOutEventAttributes workflowExecutionTimedOutEventAttributes
  80:  optional DecisionTaskScheduledEventAttributes decisionTaskScheduledEventAttributes
  90:  optional DecisionTaskStartedEventAttributes decisionTaskStartedEventAttributes
  100: optional DecisionTaskCompletedEventAttributes decisionTaskCompletedEventAttributes
  110: optional DecisionTaskTimedOutEventAttributes decisionTaskTimedOutEventAttributes
  120: optional DecisionTaskFailedEventAttributes decisionTaskFailedEventAttributes
  130: optional ActivityTaskScheduledEventAttributes activityTaskScheduledEventAttributes
  140: optional ActivityTaskStartedEventAttributes activityTaskStartedEventAttributes
  150: optional ActivityTaskCompletedEventAttributes activityTaskCompletedEventAttributes
  160: optional ActivityTaskFailedEventAttributes activityTaskFailedEventAttributes
  170: optional ActivityTaskTimedOutEventAttributes activityTaskTimedOutEventAttributes
  180: optional TimerStartedEventAttributes timerStartedEventAttributes
  190: optional TimerFiredEventAttributes timerFiredEventAttributes
  200: optional ActivityTaskCancelRequestedEventAttributes activityTaskCancelRequestedEventAttributes
  210: optional RequestCancelActivityTaskFailedEventAttributes requestCancelActivityTaskFailedEventAttributes
  220: optional ActivityTaskCanceledEventAttributes activityTaskCanceledEventAttributes
  230: optional TimerCanceledEventAttributes timerCanceledEventAttributes
  240: optional CancelTimerFailedEventAttributes cancelTimerFailedEventAttributes
  250: optional MarkerRecordedEventAttributes markerRecordedEventAttributes
  260: optional WorkflowExecutionSignaledEventAttributes workflowExecutionSignaledEventAttributes
  270: optional WorkflowExecutionTerminatedEventAttributes workflowExecutionTerminatedEventAttributes
  280: optional WorkflowExecutionCancelRequestedEventAttributes workflowExecutionCancelRequestedEventAttributes
  290: optional WorkflowExecutionCanceledEventAttributes workflowExecutionCanceledEventAttributes
  300: optional RequestCancelExternalWorkflowExecutionInitiatedEventAttributes requestCancelExternalWorkflowExecutionInitiatedEventAttributes
  310: optional RequestCancelExternalWorkflowExecutionFailedEventAttributes requestCancelExternalWorkflowExecutionFailedEventAttributes
  320: optional ExternalWorkflowExecutionCancelRequestedEventAttributes externalWorkflowExecutionCancelRequestedEventAttributes
  330: optional WorkflowExecutionContinuedAsNewEventAttributes workflowExecutionContinuedAsNewEventAttributes
  340: optional StartChildWorkflowExecutionInitiatedEventAttributes startChildWorkflowExecutionInitiatedEventAttributes
  350: optional StartChildWorkflowExecutionFailedEventAttributes startChildWorkflowExecutionFailedEventAttributes
  360: optional ChildWorkflowExecutionStartedEventAttributes childWorkflowExecutionStartedEventAttributes
  370: optional ChildWorkflowExecutionCompletedEventAttributes childWorkflowExecutionCompletedEventAttributes
  380: optional ChildWorkflowExecutionFailedEventAttributes childWorkflowExecutionFailedEventAttributes
  390: optional ChildWorkflowExecutionCanceledEventAttributes childWorkflowExecutionCanceledEventAttributes
  400: optional ChildWorkflowExecutionTimedOutEventAttributes childWorkflowExecutionTimedOutEventAttributes
  410: optional ChildWorkflowExecutionTerminatedEventAttributes childWorkflowExecutionTerminatedEventAttributes
}

struct History {
  10: optional list<HistoryEvent> events
}

struct WorkflowExecutionFilter {
  10: optional string workflowId
}

struct WorkflowTypeFilter {
  10: optional string name
}

struct StartTimeFilter {
  10: optional i64 (js.type = "Long") earliestTime
  20: optional i64 (js.type = "Long") latestTime
}

struct DomainInfo {
  10: optional string name
  20: optional DomainStatus status
  30: optional string description
  40: optional string ownerEmail
}

struct DomainConfiguration {
  10: optional i32 workflowExecutionRetentionPeriodInDays
  20: optional bool emitMetric
}

struct UpdateDomainInfo {
  10: optional string description
  20: optional string ownerEmail
}

struct RegisterDomainRequest {
  10: optional string name
  20: optional string description
  30: optional string ownerEmail
  40: optional i32 workflowExecutionRetentionPeriodInDays
  50: optional bool emitMetric
}

struct DescribeDomainRequest {
 10: optional string name
}

struct DescribeDomainResponse {
  10: optional DomainInfo domainInfo
  20: optional DomainConfiguration configuration
}

struct UpdateDomainRequest {
 10: optional string name
 20: optional UpdateDomainInfo updatedInfo
 30: optional DomainConfiguration configuration
}

struct UpdateDomainResponse {
  10: optional DomainInfo domainInfo
  20: optional DomainConfiguration configuration
}

struct DeprecateDomainRequest {
 10: optional string name
}

struct StartWorkflowExecutionRequest {
  10: optional string domain
  20: optional string workflowId
  30: optional WorkflowType workflowType
  40: optional TaskList taskList
  50: optional binary input
  60: optional i32 executionStartToCloseTimeoutSeconds
  70: optional i32 taskStartToCloseTimeoutSeconds
  80: optional string identity
  90: optional string requestId
}

struct StartWorkflowExecutionResponse {
  10: optional string runId
}

struct PollForDecisionTaskRequest {
  10: optional string domain
  20: optional TaskList taskList
  30: optional string identity
}

struct PollForDecisionTaskResponse {
  10: optional binary taskToken
  20: optional WorkflowExecution workflowExecution
  30: optional WorkflowType workflowType
  40: optional i64 (js.type = "Long") previousStartedEventId
  50: optional i64 (js.type = "Long") startedEventId
  60: optional History history
  70: optional binary nextPageToken
}

struct RespondDecisionTaskCompletedRequest {
  10: optional binary taskToken
  20: optional list<Decision> decisions
  30: optional binary executionContext
  40: optional string identity
}

struct PollForActivityTaskRequest {
  10: optional string domain
  20: optional TaskList taskList
  30: optional string identity
}

struct PollForActivityTaskResponse {
  10:  optional binary taskToken
  20:  optional WorkflowExecution workflowExecution
  30:  optional string activityId
  40:  optional ActivityType activityType
  50:  optional binary input
  60:  optional i64 (js.type = "Long") startedEventId
  70:  optional i64 (js.type = "Long") scheduledTimestamp
  80:  optional i32 scheduleToCloseTimeoutSeconds
  90:  optional i64 (js.type = "Long") startedTimestamp
  100: optional i32 startToCloseTimeoutSeconds
  110: optional i32 heartbeatTimeoutSeconds
}

struct RecordActivityTaskHeartbeatRequest {
  10: optional binary taskToken
  20: optional binary details
  30: optional string identity
}

struct RecordActivityTaskHeartbeatResponse {
  10: optional bool cancelRequested
}

struct RespondActivityTaskCompletedRequest {
  10: optional binary taskToken
  20: optional binary result
  30: optional string identity
}

struct RespondActivityTaskFailedRequest {
  10: optional binary taskToken
  20: optional string reason
  30: optional binary details
  40: optional string identity
}

struct RespondActivityTaskCanceledRequest {
  10: optional binary taskToken
  20: optional binary details
  30: optional string identity
}

struct RequestCancelWorkflowExecutionRequest {
  10: optional string domain
  20: optional WorkflowExecution workflowExecution
  30: optional string identity
}

struct GetWorkflowExecutionHistoryRequest {
  10: optional string domain
  20: optional WorkflowExecution execution
  30: optional i32 maximumPageSize
  40: optional binary nextPageToken
}

struct GetWorkflowExecutionHistoryResponse {
  10: optional History history
  20: optional binary nextPageToken
}

struct SignalWorkflowExecutionRequest {
  10: optional string domain
  20: optional WorkflowExecution workflowExecution
  30: optional string signalName
  40: optional binary input
  50: optional string identity
}

struct TerminateWorkflowExecutionRequest {
  10: optional string domain
  20: optional WorkflowExecution workflowExecution
  30: optional string reason
  40: optional binary details
  50: optional string identity
}

struct ListOpenWorkflowExecutionsRequest {
  10: optional string domain
  20: optional i32 maximumPageSize
  30: optional binary nextPageToken
  40: optional StartTimeFilter StartTimeFilter
  50: optional WorkflowExecutionFilter executionFilter
  60: optional WorkflowTypeFilter typeFilter
}

struct ListOpenWorkflowExecutionsResponse {
  10: optional list<WorkflowExecutionInfo> executions
  20: optional binary nextPageToken
}

struct ListClosedWorkflowExecutionsRequest {
  10: optional string domain
  20: optional i32 maximumPageSize
  30: optional binary nextPageToken
  40: optional StartTimeFilter StartTimeFilter
  50: optional WorkflowExecutionFilter executionFilter
  60: optional WorkflowTypeFilter typeFilter
  70: optional WorkflowExecutionCloseStatus statusFilter
}

struct ListClosedWorkflowExecutionsResponse {
  10: optional list<WorkflowExecutionInfo> executions
  20: optional binary nextPageToken
}
