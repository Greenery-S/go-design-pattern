package state

import (
	"fmt"
	"math/rand"
)

// 模拟一个审批工单的流程
// 1. 发起工单, 状态为待确认, 由发起人确认
// 2. 确认工单, 状态为待审批, 由审批人审批
//   2.1 审批通过, 状态为已通过
//   2.2 审批不通过, 状态为已拒绝
// 3. 审批通过后, 状态为已通过, 由执行人执行
// 4. 执行完成后, 状态为已完成
//  4.1 执行成功后, 状态为已完成
//  4.2 执行失败后, 状态为已失败
// 任何一步, 都可以取消工单, 状态为已取消

// 工单状态: 待确认, 待审批, 已通过, 已拒绝, 已完成, 已失败, 已取消

const (
	canceled = iota
	toConfirm
	toApprove
	approved
	rejected
	toExecute
	finished
	failed
)

// State 定义一个接口, 用于定义工单的状态
type State interface {
	Current() int
	Cancel(*Workflow)
	Next(*Workflow)
	Rollback(*Workflow)
}

// Workflow 定义一个工作流, 用于保存当前的状态
type Workflow struct {
	currentState State
}

// NewWorkflow 创建一个新的工作流
func NewWorkflow() *Workflow {
	return &Workflow{
		currentState: getWorkflowState(toConfirm),
	}
}

// Current 用于获取当前的状态
func (w *Workflow) Current() int {
	return w.currentState.Current()
}

// Cancel 用于取消工单
func (w *Workflow) Cancel() {
	fmt.Println("取消工单")
	w.currentState.Cancel(w)
}

// Next 用于获取下一个状态
func (w *Workflow) Next() {
	fmt.Println("进入下一个状态")
	w.currentState.Next(w)
}

// Rollback 用于回滚到上一个状态
func (w *Workflow) Rollback() {
	fmt.Println("回滚到上一个状态")
	w.currentState.Rollback(w)
}

func (w *Workflow) Print() string {
	fmt.Println("当前状态为: ", PrintState(w.Current()))
	return PrintState(w.Current())
}

// 用于获取当前的状态
func getWorkflowState(state int) State {
	switch state {
	case toConfirm:
		return &ToConfirm{}
	case toApprove:
		return &ToApprove{}
	case approved:
		return &Approved{}
	case rejected:
		return &Rejected{}
	case toExecute:
		return &ToExecute{}
	case finished:
		return &Finished{}
	case failed:
		return &Failed{}
	default:
		return &Canceled{}
	}
}

func PrintState(stat int) string {
	switch stat {
	case toConfirm:
		return "待确认"
	case toApprove:
		return "待审批"
	case approved:
		return "已通过"
	case rejected:
		return "已拒绝"
	case toExecute:
		return "待执行"
	case finished:
		return "已完成"
	case failed:
		return "已失败"
	default:
		return "已取消"
	}
}

func coinFlip() bool {
	if rand.Intn(2) == 0 {
		return true
	}
	return false
}

//// * 待确认状态

type ToConfirm struct{}

func (*ToConfirm) Current() int {
	return toConfirm
}

func (*ToConfirm) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*ToConfirm) Next(w *Workflow) {
	w.currentState = getWorkflowState(toApprove)
}

func (*ToConfirm) Rollback(_ *Workflow) {
	fmt.Println("当前状态为初始状态(待确认), 无法回滚")
}

//// * 待审批状态

type ToApprove struct{}

func (*ToApprove) Current() int {
	return toApprove
}

func (*ToApprove) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*ToApprove) Next(w *Workflow) {
	// 等待审批人审批
	// 这里模拟一个随机的审批结果
	if coinFlip() {
		w.currentState = getWorkflowState(approved)
	} else {
		w.currentState = getWorkflowState(rejected)
	}
}

func (*ToApprove) Rollback(w *Workflow) {
	w.currentState = getWorkflowState(toConfirm)
}

//// * 已通过状态

type Approved struct{}

func (*Approved) Current() int {
	return approved
}

func (*Approved) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*Approved) Next(w *Workflow) {
	w.currentState = getWorkflowState(toExecute)
}

func (*Approved) Rollback(w *Workflow) {
	w.currentState = getWorkflowState(toApprove)
}

//// * 已拒绝状态

type Rejected struct{}

func (*Rejected) Current() int {
	return rejected
}

func (*Rejected) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*Rejected) Next(_ *Workflow) {
	fmt.Println("当前状态为终态(已拒绝), 无法继续审批")
}

func (*Rejected) Rollback(w *Workflow) {
	w.currentState = getWorkflowState(toApprove)
}

//// * 待执行状态

type ToExecute struct{}

func (*ToExecute) Current() int {
	return toExecute
}

func (*ToExecute) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*ToExecute) Next(w *Workflow) {
	// 等待执行人执行
	// 这里模拟一个随机的执行结果
	if coinFlip() {
		w.currentState = getWorkflowState(finished)
	} else {
		w.currentState = getWorkflowState(failed)
	}
}

func (*ToExecute) Rollback(w *Workflow) {
	w.currentState = getWorkflowState(approved)
}

//// * 已完成状态

type Finished struct{}

func (*Finished) Current() int {
	return finished
}

func (*Finished) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*Finished) Next(_ *Workflow) {
	fmt.Println("当前状态为终态(已完成), 无法继续执行")
}

func (*Finished) Rollback(w *Workflow) {
	w.currentState = getWorkflowState(toExecute)
}

//// * 已取消状态

type Canceled struct{}

func (*Canceled) Current() int {
	return canceled
}

func (*Canceled) Cancel(_ *Workflow) {
	fmt.Println("当前状态为终态(已取消), 无法继续取消")
}

func (*Canceled) Next(_ *Workflow) {
	fmt.Println("当前状态为终态(已取消), 无法继续执行")
}

func (*Canceled) Rollback(_ *Workflow) {
	fmt.Println("当前状态为终态(已取消), 无法回滚")
}

//// * 已失败状态

type Failed struct{}

func (*Failed) Current() int {
	return failed
}

func (*Failed) Cancel(w *Workflow) {
	w.currentState = getWorkflowState(canceled)
}

func (*Failed) Next(_ *Workflow) {
	fmt.Println("当前状态为终态(已失败), 无法继续执行")
}

func (*Failed) Rollback(w *Workflow) {
	w.currentState = getWorkflowState(toExecute)
}
