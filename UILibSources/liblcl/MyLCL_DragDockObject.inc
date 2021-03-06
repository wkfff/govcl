
//----------------------------------------
// 代码由GenlibLcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

function DragDockObject_Create(AOwner: TControl): TDragDockObject; extdecl;
begin
  Result :=  TDragDockObject.Create(AOwner);
end;

procedure DragDockObject_Free(AObj: TDragDockObject); extdecl;
begin
  AObj.Free;
end;

procedure DragDockObject_Assign(AObj: TDragDockObject; Source: TDragObject); extdecl;
begin
  AObj.Assign(Source);
end;

procedure DragDockObject_HideDragImage(AObj: TDragDockObject); extdecl;
begin
  AObj.HideDragImage;
end;

procedure DragDockObject_ShowDragImage(AObj: TDragDockObject); extdecl;
begin
  AObj.ShowDragImage;
end;

function DragDockObject_ClassType(AObj: TDragDockObject): TClass; extdecl;
begin
  Result :=  AObj.ClassType;
end;

function DragDockObject_ClassName(AObj: TDragDockObject): PChar; extdecl;
begin
  Result :=  ToPChar(AObj.ClassName);
end;

function DragDockObject_InstanceSize(AObj: TDragDockObject): Integer; extdecl;
begin
  Result :=  AObj.InstanceSize;
end;

function DragDockObject_InheritsFrom(AObj: TDragDockObject; AClass: TClass): LongBool; extdecl;
begin
  Result :=  AObj.InheritsFrom(AClass);
end;

function DragDockObject_Equals(AObj: TDragDockObject; Obj: TObject): LongBool; extdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function DragDockObject_GetHashCode(AObj: TDragDockObject): Integer; extdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function DragDockObject_ToString(AObj: TDragDockObject): PChar; extdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

procedure DragDockObject_GetDockRect(AObj: TDragDockObject; var Result: TRect); extdecl;
begin
  Result :=  AObj.DockRect;
end;

procedure DragDockObject_SetDockRect(AObj: TDragDockObject; var AValue: TRect); extdecl;
begin
  AObj.DockRect := AValue;
end;

function DragDockObject_GetDropAlign(AObj: TDragDockObject): TAlign; extdecl;
begin
  Result :=  AObj.DropAlign;
end;

function DragDockObject_GetDropOnControl(AObj: TDragDockObject): TControl; extdecl;
begin
  Result :=  AObj.DropOnControl;
end;

procedure DragDockObject_GetEraseDockRect(AObj: TDragDockObject; var Result: TRect); extdecl;
begin
  Result :=  AObj.EraseDockRect;
end;

procedure DragDockObject_SetEraseDockRect(AObj: TDragDockObject; var AValue: TRect); extdecl;
begin
  AObj.EraseDockRect := AValue;
end;

function DragDockObject_GetFloating(AObj: TDragDockObject): LongBool; extdecl;
begin
  Result :=  AObj.Floating;
end;

procedure DragDockObject_SetFloating(AObj: TDragDockObject; AValue: LongBool); extdecl;
begin
  AObj.Floating := AValue;
end;

function DragDockObject_GetControl(AObj: TDragDockObject): TControl; extdecl;
begin
  Result :=  AObj.Control;
end;

procedure DragDockObject_SetControl(AObj: TDragDockObject; AValue: TControl); extdecl;
begin
  AObj.Control := AValue;
end;

function DragDockObject_GetAlwaysShowDragImages(AObj: TDragDockObject): LongBool; extdecl;
begin
  Result :=  AObj.AlwaysShowDragImages;
end;

procedure DragDockObject_SetAlwaysShowDragImages(AObj: TDragDockObject; AValue: LongBool); extdecl;
begin
  AObj.AlwaysShowDragImages := AValue;
end;

procedure DragDockObject_GetDragPos(AObj: TDragDockObject; var Result: TPoint); extdecl;
begin
  Result :=  AObj.DragPos;
end;

procedure DragDockObject_SetDragPos(AObj: TDragDockObject; var AValue: TPoint); extdecl;
begin
  AObj.DragPos := AValue;
end;

function DragDockObject_GetDragTarget(AObj: TDragDockObject): Pointer; extdecl;
begin
  Result :=  AObj.DragTarget;
end;

procedure DragDockObject_SetDragTarget(AObj: TDragDockObject; AValue: Pointer); extdecl;
begin
  AObj.DragTarget := AValue;
end;

procedure DragDockObject_GetDragTargetPos(AObj: TDragDockObject; var Result: TPoint); extdecl;
begin
  Result :=  AObj.DragTargetPos;
end;

procedure DragDockObject_SetDragTargetPos(AObj: TDragDockObject; var AValue: TPoint); extdecl;
begin
  AObj.DragTargetPos := AValue;
end;

function DragDockObject_GetDropped(AObj: TDragDockObject): LongBool; extdecl;
begin
  Result :=  AObj.Dropped;
end;

function DragDockObject_StaticClassType: TClass; extdecl;
begin
  Result := TDragDockObject;
end;

exports
  DragDockObject_Create,
  DragDockObject_Free,
  DragDockObject_Assign,
  DragDockObject_HideDragImage,
  DragDockObject_ShowDragImage,
  DragDockObject_ClassType,
  DragDockObject_ClassName,
  DragDockObject_InstanceSize,
  DragDockObject_InheritsFrom,
  DragDockObject_Equals,
  DragDockObject_GetHashCode,
  DragDockObject_ToString,
  DragDockObject_GetDockRect,
  DragDockObject_SetDockRect,
  DragDockObject_GetDropAlign,
  DragDockObject_GetDropOnControl,
  DragDockObject_GetEraseDockRect,
  DragDockObject_SetEraseDockRect,
  DragDockObject_GetFloating,
  DragDockObject_SetFloating,
  DragDockObject_GetControl,
  DragDockObject_SetControl,
  DragDockObject_GetAlwaysShowDragImages,
  DragDockObject_SetAlwaysShowDragImages,
  DragDockObject_GetDragPos,
  DragDockObject_SetDragPos,
  DragDockObject_GetDragTarget,
  DragDockObject_SetDragTarget,
  DragDockObject_GetDragTargetPos,
  DragDockObject_SetDragTargetPos,
  DragDockObject_GetDropped,
  DragDockObject_StaticClassType;

