package rtda

import rtc "jvmgo/jvm/rtda/class"

// stack frame
type Frame struct {
    lower           *Frame // stack is implemented as linked list
    thread          *Thread
    method          *rtc.Method
    localVars       *LocalVars
    operandStack    *OperandStack
    nextPC          int // the next instruction after the call
    onPopAction     func()
}

func newFrame(thread *Thread, method *rtc.Method) (*Frame) {
    frame := &Frame{}
    frame.thread = thread
    frame.method = method
    frame.localVars = newLocalVars(method.MaxLocals())
    frame.operandStack = newOperandStack(method.MaxStack())
    return frame
}

// getters & setters
func (self *Frame) Thread() (*Thread) {
    return self.thread
}
func (self *Frame) NextPC() (int) {
    return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
    self.nextPC = nextPC
}
func (self *Frame) LocalVars() (*LocalVars) {
    return self.localVars
}
func (self *Frame) OperandStack() (*OperandStack) {
    return self.operandStack
}
func (self *Frame) Method() (*rtc.Method) {
    return self.method
}
func (self *Frame) SetOnPopAction(f func()) {
    self.onPopAction = f
}
