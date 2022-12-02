/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	infrav1beta2 "sigs.k8s.io/cluster-api-provider-ibmcloud/api/v1beta2"
)

func (src *IBMPowerVSCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSCluster)

	return Convert_v1beta1_IBMPowerVSCluster_To_v1beta2_IBMPowerVSCluster(src, dst, nil)
}

func (dst *IBMPowerVSCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSCluster)

	return Convert_v1beta2_IBMPowerVSCluster_To_v1beta1_IBMPowerVSCluster(src, dst, nil)
}

func (src *IBMPowerVSClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSClusterList)

	return Convert_v1beta1_IBMPowerVSClusterList_To_v1beta2_IBMPowerVSClusterList(src, dst, nil)
}

func (dst *IBMPowerVSClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSClusterList)

	return Convert_v1beta2_IBMPowerVSClusterList_To_v1beta1_IBMPowerVSClusterList(src, dst, nil)
}

func (src *IBMPowerVSClusterTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSClusterTemplate)

	return Convert_v1beta1_IBMPowerVSClusterTemplate_To_v1beta2_IBMPowerVSClusterTemplate(src, dst, nil)
}

func (dst *IBMPowerVSClusterTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSClusterTemplate)

	return Convert_v1beta2_IBMPowerVSClusterTemplate_To_v1beta1_IBMPowerVSClusterTemplate(src, dst, nil)
}

func (src *IBMPowerVSClusterTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSClusterTemplateList)

	return Convert_v1beta1_IBMPowerVSClusterTemplateList_To_v1beta2_IBMPowerVSClusterTemplateList(src, dst, nil)
}

func (dst *IBMPowerVSClusterTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSClusterTemplateList)

	return Convert_v1beta2_IBMPowerVSClusterTemplateList_To_v1beta1_IBMPowerVSClusterTemplateList(src, dst, nil)
}

func (src *IBMPowerVSMachine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSMachine)

	return Convert_v1beta1_IBMPowerVSMachine_To_v1beta2_IBMPowerVSMachine(src, dst, nil)
}

func (dst *IBMPowerVSMachine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSMachine)

	return Convert_v1beta2_IBMPowerVSMachine_To_v1beta1_IBMPowerVSMachine(src, dst, nil)
}

func (src *IBMPowerVSMachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSMachineList)

	return Convert_v1beta1_IBMPowerVSMachineList_To_v1beta2_IBMPowerVSMachineList(src, dst, nil)
}

func (dst *IBMPowerVSMachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSMachineList)

	return Convert_v1beta2_IBMPowerVSMachineList_To_v1beta1_IBMPowerVSMachineList(src, dst, nil)
}

func (src *IBMPowerVSMachineTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSMachineTemplate)

	return Convert_v1beta1_IBMPowerVSMachineTemplate_To_v1beta2_IBMPowerVSMachineTemplate(src, dst, nil)
}

func (dst *IBMPowerVSMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSMachineTemplate)

	return Convert_v1beta2_IBMPowerVSMachineTemplate_To_v1beta1_IBMPowerVSMachineTemplate(src, dst, nil)
}

func (src *IBMPowerVSMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSMachineTemplateList)

	return Convert_v1beta1_IBMPowerVSMachineTemplateList_To_v1beta2_IBMPowerVSMachineTemplateList(src, dst, nil)
}

func (dst *IBMPowerVSMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSMachineTemplateList)

	return Convert_v1beta2_IBMPowerVSMachineTemplateList_To_v1beta1_IBMPowerVSMachineTemplateList(src, dst, nil)
}

func (src *IBMPowerVSImage) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSImage)

	return Convert_v1beta1_IBMPowerVSImage_To_v1beta2_IBMPowerVSImage(src, dst, nil)
}

func (dst *IBMPowerVSImage) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSImage)

	return Convert_v1beta2_IBMPowerVSImage_To_v1beta1_IBMPowerVSImage(src, dst, nil)
}

func (src *IBMPowerVSImageList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta2.IBMPowerVSImageList)

	return Convert_v1beta1_IBMPowerVSImageList_To_v1beta2_IBMPowerVSImageList(src, dst, nil)
}

func (dst *IBMPowerVSImageList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta2.IBMPowerVSImageList)

	return Convert_v1beta2_IBMPowerVSImageList_To_v1beta1_IBMPowerVSImageList(src, dst, nil)
}