/*
Copyright 2023.

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

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields that is added must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob, any “inputs” to our controller go here
type CronJobSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Represents schedule in Cron format
	//+kubebuilder:validation:MinLength=0
	//👆this indicates that the minimum string length of the Schedule field should be 0
	Schedule string `json:"schedule"`

	//This is an optional field. It tells the contoller to start the job within the deadline seconds if it misses the
	//scheduled time. Missed job executions are counted as failed.
	//+kubebuilder:validation:Minium=0
	//👆this indicates that the minimum number of seconds that the StartingDeadlineSeconds can accept should be 0
	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	//This is an optional field. It tells the controller how to treat the execution of Jobs.
	//Refer ConcurrencyPolicy type below
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	//This is an optional field. It tells the controller to suspend subsequent executions. Does not apply to
	//already executed job. Defaults to false.
	// +optional
	Suspend *bool `json:"suspend,omitempty"`

	//This specfies the job that will be created when executing a CronJob
	//+kubebuilder:validation:Minium=0
	JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`

	// The number of successful finished jobs to retain.
	//+kubebuilder:validation:Minium=0
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// The number of failed finished jobs to retain.
	// +optional
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// ConcurrencyPolicy supports 3 types:
// AllowConcurrent: allows concurrent running of CronJons
// ForbidConcurrent: forbids concurrent running of CronJons
// ReplaceConcurrent: cancels currently running CronJon and replaces it with new one.
// Default is AllowConcurrent
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
// 👆this marker tells the controller-tools this field acceots either of the abpve 3 values
type ConcurrencyPolicy string

const (
	AllowConcurrent ConcurrencyPolicy = "Allow"

	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// CronJobStatus defines the observed state of CronJob contains any information we want users or other controllers to be able to easily obtain
type CronJobStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	// Returns slice of active jobs running at the moment
	// +optional
	ActiveJobs []corev1.ObjectReference `json:"activeJobs,omitempty"`

	// Returns last time when the job was successfully scheduled
	// +optional
	LastSuccessScheduleTime *metav1.Time `json:"lastSuccessScheduleTime,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
