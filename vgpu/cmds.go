// Copyright (c) 2022, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is initially adapted from https://github.com/vulkan-go/asche
// Copyright © 2017 Maxim Kupriianov <max@kc.vc>, under the MIT License

package vgpu

import vk "github.com/vulkan-go/vulkan"

// CmdPool is a command pool and buffer
type CmdPool struct {
	Pool vk.CommandPool
	Buff vk.CommandBuffer
}

// ConfigTransient configures the pool for transient command buffers,
// which are best used for random functions, such as memory copying.
// Use SubmitWaitFree logic.
func (cp *CmdPool) ConfigTransient(dv *Device) {
	var cmdPool vk.CommandPool
	ret := vk.CreateCommandPool(dv.Device, &vk.CommandPoolCreateInfo{
		SType:            vk.StructureTypeCommandPoolCreateInfo,
		QueueFamilyIndex: dv.QueueIndex,
		Flags:            vk.CommandPoolCreateFlags(vk.CommandPoolCreateTransientBit),
	}, nil, &cmdPool)
	IfPanic(NewError(ret))
	cp.Pool = cmdPool
}

// ConfigResettable configures the pool for persistent,
// resettable command buffers, used for rendering commands.
func (cp *CmdPool) ConfigResettable(dv *Device) {
	var cmdPool vk.CommandPool
	ret := vk.CreateCommandPool(dv.Device, &vk.CommandPoolCreateInfo{
		SType:            vk.StructureTypeCommandPoolCreateInfo,
		QueueFamilyIndex: dv.QueueIndex,
		Flags:            vk.CommandPoolCreateFlags(vk.CommandPoolCreateResetCommandBufferBit),
	}, nil, &cmdPool)
	IfPanic(NewError(ret))
	cp.Pool = cmdPool
}

// NewBuffer makes a buffer in pool, setting Buff to point to it
// and also returning the buffer.
func (cp *CmdPool) NewBuffer(dv *Device) vk.CommandBuffer {
	var cmdBuff = make([]vk.CommandBuffer, 1)
	ret := vk.AllocateCommandBuffers(dv.Device, &vk.CommandBufferAllocateInfo{
		SType:              vk.StructureTypeCommandBufferAllocateInfo,
		CommandPool:        cp.Pool,
		Level:              vk.CommandBufferLevelPrimary,
		CommandBufferCount: 1,
	}, cmdBuff)
	IfPanic(NewError(ret))
	cBuff := cmdBuff[0]
	cp.Buff = cBuff
	return cBuff
}

// BeginCmd does BeginCommandBuffer on buffer
func (cp *CmdPool) BeginCmd() vk.CommandBuffer {
	CmdBegin(cp.Buff)
	return cp.Buff
}

// BeginCmdOneTime does BeginCommandBuffer with OneTimeSubmit set on buffer
func (cp *CmdPool) BeginCmdOneTime() vk.CommandBuffer {
	CmdBeginOneTime(cp.Buff)
	return cp.Buff
}

// SubmitWait does End, Submit, WaitIdle on Buffer
func (cp *CmdPool) SubmitWait(dev *Device) {
	CmdSubmitWait(cp.Buff, dev)
}

// SubmitWaitFree does End, Submit, WaitIdle, Free on Buffer
func (cp *CmdPool) SubmitWaitFree(dev *Device) {
	cp.SubmitWait(dev)
	cp.FreeBuffer(dev)
}

// EndCmd does EndCommandBuffer on buffer
func (cp *CmdPool) EndCmd() {
	CmdEnd(cp.Buff)
}

// Submit submits commands in buffer to given device queue, without
// any semaphore logic -- suitable for a WaitIdle logic.
func (cp *CmdPool) Submit(dev *Device) {
	CmdSubmit(cp.Buff, dev)
}

// Reset resets the command buffer so it is ready for recording new commands.
func (cp *CmdPool) Reset() {
	CmdReset(cp.Buff)
}

// FreeBuffer frees the current Buff buffer
func (cp *CmdPool) FreeBuffer(dev *Device) {
	cmdBu := []vk.CommandBuffer{cp.Buff}
	vk.FreeCommandBuffers(dev.Device, cp.Pool, 1, cmdBu)
	cp.Buff = nil
}

// Destroy
func (cp *CmdPool) Destroy(dev vk.Device) {
	if cp.Pool == nil {
		return
	}
	vk.DestroyCommandPool(dev, cp.Pool, nil)
	cp.Pool = nil
}

//////////////////////////////////////////////////////////////
// Command Buffer functions

// CmdBegin does BeginCommandBuffer on buffer
func CmdBegin(cmd vk.CommandBuffer) {
	ret := vk.BeginCommandBuffer(cmd, &vk.CommandBufferBeginInfo{
		SType: vk.StructureTypeCommandBufferBeginInfo,
	})
	IfPanic(NewError(ret))
}

// CmdBeginOneTime does BeginCommandBuffer with OneTimeSubmit set on buffer
func CmdBeginOneTime(cmd vk.CommandBuffer) {
	ret := vk.BeginCommandBuffer(cmd, &vk.CommandBufferBeginInfo{
		SType: vk.StructureTypeCommandBufferBeginInfo,
		Flags: vk.CommandBufferUsageFlags(vk.CommandBufferUsageOneTimeSubmitBit),
	})
	IfPanic(NewError(ret))
}

// CmdSubmit submits commands in buffer to given device queue, without
// any semaphore logic -- suitable for a WaitIdle logic.
func CmdSubmit(cmd vk.CommandBuffer, dev *Device) {
	ret := vk.QueueSubmit(dev.Queue, 1, []vk.SubmitInfo{{
		SType:              vk.StructureTypeSubmitInfo,
		CommandBufferCount: 1,
		PCommandBuffers:    []vk.CommandBuffer{cmd},
	}}, vk.NullFence)
	IfPanic(NewError(ret))
}

// CmdSubmitWait does End, Submit, WaitIdle on command Buffer
func CmdSubmitWait(cmd vk.CommandBuffer, dev *Device) {
	CmdEnd(cmd)
	CmdSubmit(cmd, dev)
	vk.QueueWaitIdle(dev.Queue)
}

// CmdEnd does EndCommandBuffer on buffer
func CmdEnd(cmd vk.CommandBuffer) {
	ret := vk.EndCommandBuffer(cmd)
	IfPanic(NewError(ret))
}

// CmdReset resets the command buffer so it is ready for recording new commands.
func CmdReset(cmd vk.CommandBuffer) {
	vk.ResetCommandBuffer(cmd, 0)
}

//////////////////////////////////////////////////////////////
// Semaphors & Fences

func NewSemaphore(dev vk.Device) vk.Semaphore {
	var sem vk.Semaphore
	ret := vk.CreateSemaphore(dev, &vk.SemaphoreCreateInfo{
		SType: vk.StructureTypeSemaphoreCreateInfo,
	}, nil, &sem)
	IfPanic(NewError(ret))
	return sem
}

func NewFence(dev vk.Device) vk.Fence {
	var fence vk.Fence
	ret := vk.CreateFence(dev, &vk.FenceCreateInfo{
		SType: vk.StructureTypeFenceCreateInfo,
		Flags: vk.FenceCreateFlags(vk.FenceCreateSignaledBit),
	}, nil, &fence)
	IfPanic(NewError(ret))
	return fence
}
