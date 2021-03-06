
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TTrackBar struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTrackBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTrackBar(owner IComponent) *TTrackBar {
    t := new(TTrackBar)
    t.instance = TrackBar_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TrackBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TrackBarFromInst(inst uintptr) *TTrackBar {
    t := new(TTrackBar)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TrackBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TrackBarFromObj(obj IObject) *TTrackBar {
    t := new(TTrackBar)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TrackBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TrackBarFromUnsafePointer(ptr unsafe.Pointer) *TTrackBar {
    t := new(TTrackBar)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTrackBar) Free() {
    if t.instance != 0 {
        TrackBar_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTrackBar) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTrackBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTrackBar) IsValid() bool {
    return t.instance != 0
}

// TTrackBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTrackBarClass() TClass {
    return TrackBar_StaticClassType()
}

// SetTick
func (t *TTrackBar) SetTick(Value int32) {
    TrackBar_SetTick(t.instance, Value)
}

// CanFocus
// CN: 是否可以获得焦点。
// EN: .
func (t *TTrackBar) CanFocus() bool {
    return TrackBar_CanFocus(t.instance)
}

// ContainsControl
// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (t *TTrackBar) ContainsControl(Control IControl) bool {
    return TrackBar_ContainsControl(t.instance, CheckPtr(Control))
}

// ControlAtPos
// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (t *TTrackBar) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(TrackBar_ControlAtPos(t.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// DisableAlign
// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (t *TTrackBar) DisableAlign() {
    TrackBar_DisableAlign(t.instance)
}

// EnableAlign
// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (t *TTrackBar) EnableAlign() {
    TrackBar_EnableAlign(t.instance)
}

// FindChildControl
// CN: 查找子控件。
// EN: Find sub controls.
func (t *TTrackBar) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(TrackBar_FindChildControl(t.instance, ControlName))
}

// FlipChildren
func (t *TTrackBar) FlipChildren(AllLevels bool) {
    TrackBar_FlipChildren(t.instance, AllLevels)
}

// Focused
// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (t *TTrackBar) Focused() bool {
    return TrackBar_Focused(t.instance)
}

// HandleAllocated
// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (t *TTrackBar) HandleAllocated() bool {
    return TrackBar_HandleAllocated(t.instance)
}

// InsertControl
// CN: 插入一个控件。
// EN: Insert a control.
func (t *TTrackBar) InsertControl(AControl IControl) {
    TrackBar_InsertControl(t.instance, CheckPtr(AControl))
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (t *TTrackBar) Invalidate() {
    TrackBar_Invalidate(t.instance)
}

// PaintTo
// CN: 绘画至指定DC。
// EN: Painting to the specified DC.
func (t *TTrackBar) PaintTo(DC HDC, X int32, Y int32) {
    TrackBar_PaintTo(t.instance, DC , X , Y)
}

// RemoveControl
// CN: 移除一个控件。
// EN: Remove a control.
func (t *TTrackBar) RemoveControl(AControl IControl) {
    TrackBar_RemoveControl(t.instance, CheckPtr(AControl))
}

// Realign
// CN: 重新对齐。
// EN: Realign.
func (t *TTrackBar) Realign() {
    TrackBar_Realign(t.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (t *TTrackBar) Repaint() {
    TrackBar_Repaint(t.instance)
}

// ScaleBy
// CN: 按比例缩放。
// EN: Scale by.
func (t *TTrackBar) ScaleBy(M int32, D int32) {
    TrackBar_ScaleBy(t.instance, M , D)
}

// ScrollBy
// CN: 滚动至指定位置。
// EN: Scroll by.
func (t *TTrackBar) ScrollBy(DeltaX int32, DeltaY int32) {
    TrackBar_ScrollBy(t.instance, DeltaX , DeltaY)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (t *TTrackBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    TrackBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// SetFocus
// CN: 设置控件焦点。
// EN: Set control focus.
func (t *TTrackBar) SetFocus() {
    TrackBar_SetFocus(t.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (t *TTrackBar) Update() {
    TrackBar_Update(t.instance)
}

// UpdateControlState
// CN: 更新控件状态。
// EN: Update control status.
func (t *TTrackBar) UpdateControlState() {
    TrackBar_UpdateControlState(t.instance)
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (t *TTrackBar) BringToFront() {
    TrackBar_BringToFront(t.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (t *TTrackBar) ClientToScreen(Point TPoint) TPoint {
    return TrackBar_ClientToScreen(t.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (t *TTrackBar) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return TrackBar_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (t *TTrackBar) Dragging() bool {
    return TrackBar_Dragging(t.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (t *TTrackBar) HasParent() bool {
    return TrackBar_HasParent(t.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (t *TTrackBar) Hide() {
    TrackBar_Hide(t.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (t *TTrackBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return TrackBar_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (t *TTrackBar) Refresh() {
    TrackBar_Refresh(t.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (t *TTrackBar) ScreenToClient(Point TPoint) TPoint {
    return TrackBar_ScreenToClient(t.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (t *TTrackBar) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return TrackBar_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (t *TTrackBar) SendToBack() {
    TrackBar_SendToBack(t.instance)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (t *TTrackBar) Show() {
    TrackBar_Show(t.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (t *TTrackBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    return TrackBar_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (t *TTrackBar) GetTextLen() int32 {
    return TrackBar_GetTextLen(t.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (t *TTrackBar) SetTextBuf(Buffer string) {
    TrackBar_SetTextBuf(t.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (t *TTrackBar) FindComponent(AName string) *TComponent {
    return ComponentFromInst(TrackBar_FindComponent(t.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTrackBar) GetNamePath() string {
    return TrackBar_GetNamePath(t.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTrackBar) Assign(Source IObject) {
    TrackBar_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTrackBar) DisposeOf() {
    TrackBar_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTrackBar) ClassType() TClass {
    return TrackBar_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTrackBar) ClassName() string {
    return TrackBar_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTrackBar) InstanceSize() int32 {
    return TrackBar_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTrackBar) InheritsFrom(AClass TClass) bool {
    return TrackBar_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTrackBar) Equals(Obj IObject) bool {
    return TrackBar_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTrackBar) GetHashCode() int32 {
    return TrackBar_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTrackBar) ToString() string {
    return TrackBar_ToString(t.instance)
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (t *TTrackBar) Align() TAlign {
    return TrackBar_GetAlign(t.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (t *TTrackBar) SetAlign(value TAlign) {
    TrackBar_SetAlign(t.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (t *TTrackBar) Anchors() TAnchors {
    return TrackBar_GetAnchors(t.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (t *TTrackBar) SetAnchors(value TAnchors) {
    TrackBar_SetAnchors(t.instance, value)
}

// BorderWidth
// CN: 获取边框的宽度。
// EN: .
func (t *TTrackBar) BorderWidth() int32 {
    return TrackBar_GetBorderWidth(t.instance)
}

// SetBorderWidth
// CN: 设置边框的宽度。
// EN: .
func (t *TTrackBar) SetBorderWidth(value int32) {
    TrackBar_SetBorderWidth(t.instance, value)
}

// DoubleBuffered
// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (t *TTrackBar) DoubleBuffered() bool {
    return TrackBar_GetDoubleBuffered(t.instance)
}

// SetDoubleBuffered
// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (t *TTrackBar) SetDoubleBuffered(value bool) {
    TrackBar_SetDoubleBuffered(t.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (t *TTrackBar) DragCursor() TCursor {
    return TrackBar_GetDragCursor(t.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (t *TTrackBar) SetDragCursor(value TCursor) {
    TrackBar_SetDragCursor(t.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (t *TTrackBar) DragKind() TDragKind {
    return TrackBar_GetDragKind(t.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (t *TTrackBar) SetDragKind(value TDragKind) {
    TrackBar_SetDragKind(t.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (t *TTrackBar) DragMode() TDragMode {
    return TrackBar_GetDragMode(t.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (t *TTrackBar) SetDragMode(value TDragMode) {
    TrackBar_SetDragMode(t.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TTrackBar) Enabled() bool {
    return TrackBar_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TTrackBar) SetEnabled(value bool) {
    TrackBar_SetEnabled(t.instance, value)
}

// LineSize
func (t *TTrackBar) LineSize() int32 {
    return TrackBar_GetLineSize(t.instance)
}

// SetLineSize
func (t *TTrackBar) SetLineSize(value int32) {
    TrackBar_SetLineSize(t.instance, value)
}

// Max
func (t *TTrackBar) Max() int32 {
    return TrackBar_GetMax(t.instance)
}

// SetMax
func (t *TTrackBar) SetMax(value int32) {
    TrackBar_SetMax(t.instance, value)
}

// Min
func (t *TTrackBar) Min() int32 {
    return TrackBar_GetMin(t.instance)
}

// SetMin
func (t *TTrackBar) SetMin(value int32) {
    TrackBar_SetMin(t.instance, value)
}

// Orientation
func (t *TTrackBar) Orientation() TTrackBarOrientation {
    return TrackBar_GetOrientation(t.instance)
}

// SetOrientation
func (t *TTrackBar) SetOrientation(value TTrackBarOrientation) {
    TrackBar_SetOrientation(t.instance, value)
}

// ParentCtl3D
func (t *TTrackBar) ParentCtl3D() bool {
    return TrackBar_GetParentCtl3D(t.instance)
}

// SetParentCtl3D
func (t *TTrackBar) SetParentCtl3D(value bool) {
    TrackBar_SetParentCtl3D(t.instance, value)
}

// ParentDoubleBuffered
// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (t *TTrackBar) ParentDoubleBuffered() bool {
    return TrackBar_GetParentDoubleBuffered(t.instance)
}

// SetParentDoubleBuffered
// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (t *TTrackBar) SetParentDoubleBuffered(value bool) {
    TrackBar_SetParentDoubleBuffered(t.instance, value)
}

// ParentShowHint
func (t *TTrackBar) ParentShowHint() bool {
    return TrackBar_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TTrackBar) SetParentShowHint(value bool) {
    TrackBar_SetParentShowHint(t.instance, value)
}

// PageSize
func (t *TTrackBar) PageSize() int32 {
    return TrackBar_GetPageSize(t.instance)
}

// SetPageSize
func (t *TTrackBar) SetPageSize(value int32) {
    TrackBar_SetPageSize(t.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (t *TTrackBar) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(TrackBar_GetPopupMenu(t.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (t *TTrackBar) SetPopupMenu(value IComponent) {
    TrackBar_SetPopupMenu(t.instance, CheckPtr(value))
}

// Frequency
func (t *TTrackBar) Frequency() int32 {
    return TrackBar_GetFrequency(t.instance)
}

// SetFrequency
func (t *TTrackBar) SetFrequency(value int32) {
    TrackBar_SetFrequency(t.instance, value)
}

// Position
func (t *TTrackBar) Position() int32 {
    return TrackBar_GetPosition(t.instance)
}

// SetPosition
func (t *TTrackBar) SetPosition(value int32) {
    TrackBar_SetPosition(t.instance, value)
}

// PositionToolTip
func (t *TTrackBar) PositionToolTip() TPositionToolTip {
    return TrackBar_GetPositionToolTip(t.instance)
}

// SetPositionToolTip
func (t *TTrackBar) SetPositionToolTip(value TPositionToolTip) {
    TrackBar_SetPositionToolTip(t.instance, value)
}

// SliderVisible
func (t *TTrackBar) SliderVisible() bool {
    return TrackBar_GetSliderVisible(t.instance)
}

// SetSliderVisible
func (t *TTrackBar) SetSliderVisible(value bool) {
    TrackBar_SetSliderVisible(t.instance, value)
}

// SelEnd
func (t *TTrackBar) SelEnd() int32 {
    return TrackBar_GetSelEnd(t.instance)
}

// SetSelEnd
func (t *TTrackBar) SetSelEnd(value int32) {
    TrackBar_SetSelEnd(t.instance, value)
}

// SelStart
// CN: 获取选择的启始位置。
// EN: .
func (t *TTrackBar) SelStart() int32 {
    return TrackBar_GetSelStart(t.instance)
}

// SetSelStart
// CN: 设置选择的启始位置。
// EN: .
func (t *TTrackBar) SetSelStart(value int32) {
    TrackBar_SetSelStart(t.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (t *TTrackBar) ShowHint() bool {
    return TrackBar_GetShowHint(t.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (t *TTrackBar) SetShowHint(value bool) {
    TrackBar_SetShowHint(t.instance, value)
}

// ShowSelRange
func (t *TTrackBar) ShowSelRange() bool {
    return TrackBar_GetShowSelRange(t.instance)
}

// SetShowSelRange
func (t *TTrackBar) SetShowSelRange(value bool) {
    TrackBar_SetShowSelRange(t.instance, value)
}

// TabOrder
// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (t *TTrackBar) TabOrder() TTabOrder {
    return TrackBar_GetTabOrder(t.instance)
}

// SetTabOrder
// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (t *TTrackBar) SetTabOrder(value TTabOrder) {
    TrackBar_SetTabOrder(t.instance, value)
}

// TabStop
// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (t *TTrackBar) TabStop() bool {
    return TrackBar_GetTabStop(t.instance)
}

// SetTabStop
// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (t *TTrackBar) SetTabStop(value bool) {
    TrackBar_SetTabStop(t.instance, value)
}

// ThumbLength
func (t *TTrackBar) ThumbLength() int32 {
    return TrackBar_GetThumbLength(t.instance)
}

// SetThumbLength
func (t *TTrackBar) SetThumbLength(value int32) {
    TrackBar_SetThumbLength(t.instance, value)
}

// TickMarks
func (t *TTrackBar) TickMarks() TTickMark {
    return TrackBar_GetTickMarks(t.instance)
}

// SetTickMarks
func (t *TTrackBar) SetTickMarks(value TTickMark) {
    TrackBar_SetTickMarks(t.instance, value)
}

// TickStyle
func (t *TTrackBar) TickStyle() TTickStyle {
    return TrackBar_GetTickStyle(t.instance)
}

// SetTickStyle
func (t *TTrackBar) SetTickStyle(value TTickStyle) {
    TrackBar_SetTickStyle(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TTrackBar) Visible() bool {
    return TrackBar_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TTrackBar) SetVisible(value bool) {
    TrackBar_SetVisible(t.instance, value)
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (t *TTrackBar) StyleElements() TStyleElements {
    return TrackBar_GetStyleElements(t.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (t *TTrackBar) SetStyleElements(value TStyleElements) {
    TrackBar_SetStyleElements(t.instance, value)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (t *TTrackBar) SetOnContextPopup(fn TContextPopupEvent) {
    TrackBar_SetOnContextPopup(t.instance, fn)
}

// SetOnChange
// CN: 设置改变事件。
// EN: Set changed event.
func (t *TTrackBar) SetOnChange(fn TNotifyEvent) {
    TrackBar_SetOnChange(t.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (t *TTrackBar) SetOnDragDrop(fn TDragDropEvent) {
    TrackBar_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (t *TTrackBar) SetOnDragOver(fn TDragOverEvent) {
    TrackBar_SetOnDragOver(t.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (t *TTrackBar) SetOnEndDock(fn TEndDragEvent) {
    TrackBar_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (t *TTrackBar) SetOnEndDrag(fn TEndDragEvent) {
    TrackBar_SetOnEndDrag(t.instance, fn)
}

// SetOnEnter
// CN: 设置焦点进入。
// EN: Set Focus entry.
func (t *TTrackBar) SetOnEnter(fn TNotifyEvent) {
    TrackBar_SetOnEnter(t.instance, fn)
}

// SetOnExit
// CN: 设置焦点退出。
// EN: Set Focus exit.
func (t *TTrackBar) SetOnExit(fn TNotifyEvent) {
    TrackBar_SetOnExit(t.instance, fn)
}

// SetOnGesture
func (t *TTrackBar) SetOnGesture(fn TGestureEvent) {
    TrackBar_SetOnGesture(t.instance, fn)
}

// SetOnKeyDown
// CN: 设置键盘按键按下事件。
// EN: Set Keyboard button press event.
func (t *TTrackBar) SetOnKeyDown(fn TKeyEvent) {
    TrackBar_SetOnKeyDown(t.instance, fn)
}

// SetOnKeyPress
func (t *TTrackBar) SetOnKeyPress(fn TKeyPressEvent) {
    TrackBar_SetOnKeyPress(t.instance, fn)
}

// SetOnKeyUp
// CN: 设置键盘按键抬起事件。
// EN: Set Keyboard button lift event.
func (t *TTrackBar) SetOnKeyUp(fn TKeyEvent) {
    TrackBar_SetOnKeyUp(t.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (t *TTrackBar) SetOnStartDock(fn TStartDockEvent) {
    TrackBar_SetOnStartDock(t.instance, fn)
}

// DockClientCount
// CN: 获取依靠客户端总数。
// EN: .
func (t *TTrackBar) DockClientCount() int32 {
    return TrackBar_GetDockClientCount(t.instance)
}

// DockSite
// CN: 获取停靠站点。
// EN: Get Docking site.
func (t *TTrackBar) DockSite() bool {
    return TrackBar_GetDockSite(t.instance)
}

// SetDockSite
// CN: 设置停靠站点。
// EN: Set Docking site.
func (t *TTrackBar) SetDockSite(value bool) {
    TrackBar_SetDockSite(t.instance, value)
}

// AlignDisabled
// CN: 获取禁用对齐。
// EN: .
func (t *TTrackBar) AlignDisabled() bool {
    return TrackBar_GetAlignDisabled(t.instance)
}

// MouseInClient
// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (t *TTrackBar) MouseInClient() bool {
    return TrackBar_GetMouseInClient(t.instance)
}

// VisibleDockClientCount
// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (t *TTrackBar) VisibleDockClientCount() int32 {
    return TrackBar_GetVisibleDockClientCount(t.instance)
}

// Brush
// CN: 获取画刷对象。
// EN: Get Brush.
func (t *TTrackBar) Brush() *TBrush {
    return BrushFromInst(TrackBar_GetBrush(t.instance))
}

// ControlCount
// CN: 获取子控件数。
// EN: Get Number of child controls.
func (t *TTrackBar) ControlCount() int32 {
    return TrackBar_GetControlCount(t.instance)
}

// Handle
// CN: 获取控件句柄。
// EN: Get Control handle.
func (t *TTrackBar) Handle() HWND {
    return TrackBar_GetHandle(t.instance)
}

// ParentWindow
// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (t *TTrackBar) ParentWindow() HWND {
    return TrackBar_GetParentWindow(t.instance)
}

// SetParentWindow
// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (t *TTrackBar) SetParentWindow(value HWND) {
    TrackBar_SetParentWindow(t.instance, value)
}

// UseDockManager
// CN: 获取使用停靠管理。
// EN: .
func (t *TTrackBar) UseDockManager() bool {
    return TrackBar_GetUseDockManager(t.instance)
}

// SetUseDockManager
// CN: 设置使用停靠管理。
// EN: .
func (t *TTrackBar) SetUseDockManager(value bool) {
    TrackBar_SetUseDockManager(t.instance, value)
}

// Action
func (t *TTrackBar) Action() *TAction {
    return ActionFromInst(TrackBar_GetAction(t.instance))
}

// SetAction
func (t *TTrackBar) SetAction(value IComponent) {
    TrackBar_SetAction(t.instance, CheckPtr(value))
}

// BiDiMode
func (t *TTrackBar) BiDiMode() TBiDiMode {
    return TrackBar_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TTrackBar) SetBiDiMode(value TBiDiMode) {
    TrackBar_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TTrackBar) BoundsRect() TRect {
    return TrackBar_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TTrackBar) SetBoundsRect(value TRect) {
    TrackBar_SetBoundsRect(t.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (t *TTrackBar) ClientHeight() int32 {
    return TrackBar_GetClientHeight(t.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (t *TTrackBar) SetClientHeight(value int32) {
    TrackBar_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TTrackBar) ClientOrigin() TPoint {
    return TrackBar_GetClientOrigin(t.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (t *TTrackBar) ClientRect() TRect {
    return TrackBar_GetClientRect(t.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (t *TTrackBar) ClientWidth() int32 {
    return TrackBar_GetClientWidth(t.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (t *TTrackBar) SetClientWidth(value int32) {
    TrackBar_SetClientWidth(t.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (t *TTrackBar) ControlState() TControlState {
    return TrackBar_GetControlState(t.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (t *TTrackBar) SetControlState(value TControlState) {
    TrackBar_SetControlState(t.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (t *TTrackBar) ControlStyle() TControlStyle {
    return TrackBar_GetControlStyle(t.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (t *TTrackBar) SetControlStyle(value TControlStyle) {
    TrackBar_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TTrackBar) ExplicitLeft() int32 {
    return TrackBar_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TTrackBar) ExplicitTop() int32 {
    return TrackBar_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TTrackBar) ExplicitWidth() int32 {
    return TrackBar_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TTrackBar) ExplicitHeight() int32 {
    return TrackBar_GetExplicitHeight(t.instance)
}

// Floating
func (t *TTrackBar) Floating() bool {
    return TrackBar_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TTrackBar) Parent() *TWinControl {
    return WinControlFromInst(TrackBar_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TTrackBar) SetParent(value IWinControl) {
    TrackBar_SetParent(t.instance, CheckPtr(value))
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (t *TTrackBar) AlignWithMargins() bool {
    return TrackBar_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (t *TTrackBar) SetAlignWithMargins(value bool) {
    TrackBar_SetAlignWithMargins(t.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (t *TTrackBar) Left() int32 {
    return TrackBar_GetLeft(t.instance)
}

// SetLeft
// CN: 设置左边位置。
// EN: Set Left position.
func (t *TTrackBar) SetLeft(value int32) {
    TrackBar_SetLeft(t.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (t *TTrackBar) Top() int32 {
    return TrackBar_GetTop(t.instance)
}

// SetTop
// CN: 设置顶边位置。
// EN: Set Top position.
func (t *TTrackBar) SetTop(value int32) {
    TrackBar_SetTop(t.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (t *TTrackBar) Width() int32 {
    return TrackBar_GetWidth(t.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (t *TTrackBar) SetWidth(value int32) {
    TrackBar_SetWidth(t.instance, value)
}

// Height
// CN: 获取高度。
// EN: Get height.
func (t *TTrackBar) Height() int32 {
    return TrackBar_GetHeight(t.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (t *TTrackBar) SetHeight(value int32) {
    TrackBar_SetHeight(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TTrackBar) Cursor() TCursor {
    return TrackBar_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TTrackBar) SetCursor(value TCursor) {
    TrackBar_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (t *TTrackBar) Hint() string {
    return TrackBar_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (t *TTrackBar) SetHint(value string) {
    TrackBar_SetHint(t.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (t *TTrackBar) Margins() *TMargins {
    return MarginsFromInst(TrackBar_GetMargins(t.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (t *TTrackBar) SetMargins(value *TMargins) {
    TrackBar_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (t *TTrackBar) CustomHint() *TCustomHint {
    return CustomHintFromInst(TrackBar_GetCustomHint(t.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (t *TTrackBar) SetCustomHint(value IComponent) {
    TrackBar_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TTrackBar) ComponentCount() int32 {
    return TrackBar_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TTrackBar) ComponentIndex() int32 {
    return TrackBar_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TTrackBar) SetComponentIndex(value int32) {
    TrackBar_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TTrackBar) Owner() *TComponent {
    return ComponentFromInst(TrackBar_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TTrackBar) Name() string {
    return TrackBar_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TTrackBar) SetName(value string) {
    TrackBar_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TTrackBar) Tag() int {
    return TrackBar_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TTrackBar) SetTag(value int) {
    TrackBar_SetTag(t.instance, value)
}

// DockClients
// CN: 获取指定索引停靠客户端。
// EN: .
func (t *TTrackBar) DockClients(Index int32) *TControl {
    return ControlFromInst(TrackBar_GetDockClients(t.instance, Index))
}

// Controls
// CN: 获取指定索引子控件。
// EN: .
func (t *TTrackBar) Controls(Index int32) *TControl {
    return ControlFromInst(TrackBar_GetControls(t.instance, Index))
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TTrackBar) Components(AIndex int32) *TComponent {
    return ComponentFromInst(TrackBar_GetComponents(t.instance, AIndex))
}

