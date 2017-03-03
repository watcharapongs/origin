// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BinaryBuildRequestOptions, InType: reflect.TypeOf(&BinaryBuildRequestOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BinaryBuildSource, InType: reflect.TypeOf(&BinaryBuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Build, InType: reflect.TypeOf(&Build{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildConfig, InType: reflect.TypeOf(&BuildConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildConfigList, InType: reflect.TypeOf(&BuildConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildConfigSpec, InType: reflect.TypeOf(&BuildConfigSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildConfigStatus, InType: reflect.TypeOf(&BuildConfigStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildList, InType: reflect.TypeOf(&BuildList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildLog, InType: reflect.TypeOf(&BuildLog{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildLogOptions, InType: reflect.TypeOf(&BuildLogOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildOutput, InType: reflect.TypeOf(&BuildOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildPostCommitSpec, InType: reflect.TypeOf(&BuildPostCommitSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildRequest, InType: reflect.TypeOf(&BuildRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildSource, InType: reflect.TypeOf(&BuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildSpec, InType: reflect.TypeOf(&BuildSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildStatus, InType: reflect.TypeOf(&BuildStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildStatusOutput, InType: reflect.TypeOf(&BuildStatusOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildStatusOutputTo, InType: reflect.TypeOf(&BuildStatusOutputTo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildStrategy, InType: reflect.TypeOf(&BuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildTriggerCause, InType: reflect.TypeOf(&BuildTriggerCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BuildTriggerPolicy, InType: reflect.TypeOf(&BuildTriggerPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_CommonSpec, InType: reflect.TypeOf(&CommonSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_CustomBuildStrategy, InType: reflect.TypeOf(&CustomBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DockerBuildStrategy, InType: reflect.TypeOf(&DockerBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_DockerStrategyOptions, InType: reflect.TypeOf(&DockerStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GenericWebHookCause, InType: reflect.TypeOf(&GenericWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GenericWebHookEvent, InType: reflect.TypeOf(&GenericWebHookEvent{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GitBuildSource, InType: reflect.TypeOf(&GitBuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GitHubWebHookCause, InType: reflect.TypeOf(&GitHubWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GitInfo, InType: reflect.TypeOf(&GitInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_GitSourceRevision, InType: reflect.TypeOf(&GitSourceRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageChangeCause, InType: reflect.TypeOf(&ImageChangeCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageChangeTrigger, InType: reflect.TypeOf(&ImageChangeTrigger{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageLabel, InType: reflect.TypeOf(&ImageLabel{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageSource, InType: reflect.TypeOf(&ImageSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ImageSourcePath, InType: reflect.TypeOf(&ImageSourcePath{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_JenkinsPipelineBuildStrategy, InType: reflect.TypeOf(&JenkinsPipelineBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProxyConfig, InType: reflect.TypeOf(&ProxyConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SecretBuildSource, InType: reflect.TypeOf(&SecretBuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SecretSpec, InType: reflect.TypeOf(&SecretSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SourceBuildStrategy, InType: reflect.TypeOf(&SourceBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SourceControlUser, InType: reflect.TypeOf(&SourceControlUser{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_SourceRevision, InType: reflect.TypeOf(&SourceRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_WebHookTrigger, InType: reflect.TypeOf(&WebHookTrigger{})},
	)
}

func DeepCopy_v1_BinaryBuildRequestOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BinaryBuildRequestOptions)
		out := out.(*BinaryBuildRequestOptions)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		out.AsFile = in.AsFile
		out.Commit = in.Commit
		out.Message = in.Message
		out.AuthorName = in.AuthorName
		out.AuthorEmail = in.AuthorEmail
		out.CommitterName = in.CommitterName
		out.CommitterEmail = in.CommitterEmail
		return nil
	}
}

func DeepCopy_v1_BinaryBuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BinaryBuildSource)
		out := out.(*BinaryBuildSource)
		out.AsFile = in.AsFile
		return nil
	}
}

func DeepCopy_v1_Build(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Build)
		out := out.(*Build)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_BuildSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_BuildStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_BuildConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfig)
		out := out.(*BuildConfig)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_BuildConfigSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		out.Status = in.Status
		return nil
	}
}

func DeepCopy_v1_BuildConfigList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfigList)
		out := out.(*BuildConfigList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]BuildConfig, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_BuildConfig(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildConfigSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfigSpec)
		out := out.(*BuildConfigSpec)
		if in.Triggers != nil {
			in, out := &in.Triggers, &out.Triggers
			*out = make([]BuildTriggerPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_BuildTriggerPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Triggers = nil
		}
		out.RunPolicy = in.RunPolicy
		if err := DeepCopy_v1_CommonSpec(&in.CommonSpec, &out.CommonSpec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_BuildConfigStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfigStatus)
		out := out.(*BuildConfigStatus)
		out.LastVersion = in.LastVersion
		return nil
	}
}

func DeepCopy_v1_BuildList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildList)
		out := out.(*BuildList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Build, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Build(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildLog(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildLog)
		out := out.(*BuildLog)
		out.TypeMeta = in.TypeMeta
		return nil
	}
}

func DeepCopy_v1_BuildLogOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildLogOptions)
		out := out.(*BuildLogOptions)
		out.TypeMeta = in.TypeMeta
		out.Container = in.Container
		out.Follow = in.Follow
		out.Previous = in.Previous
		if in.SinceSeconds != nil {
			in, out := &in.SinceSeconds, &out.SinceSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.SinceSeconds = nil
		}
		if in.SinceTime != nil {
			in, out := &in.SinceTime, &out.SinceTime
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.SinceTime = nil
		}
		out.Timestamps = in.Timestamps
		if in.TailLines != nil {
			in, out := &in.TailLines, &out.TailLines
			*out = new(int64)
			**out = **in
		} else {
			out.TailLines = nil
		}
		if in.LimitBytes != nil {
			in, out := &in.LimitBytes, &out.LimitBytes
			*out = new(int64)
			**out = **in
		} else {
			out.LimitBytes = nil
		}
		out.NoWait = in.NoWait
		if in.Version != nil {
			in, out := &in.Version, &out.Version
			*out = new(int64)
			**out = **in
		} else {
			out.Version = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildOutput(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildOutput)
		out := out.(*BuildOutput)
		if in.To != nil {
			in, out := &in.To, &out.To
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.To = nil
		}
		if in.PushSecret != nil {
			in, out := &in.PushSecret, &out.PushSecret
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.PushSecret = nil
		}
		if in.ImageLabels != nil {
			in, out := &in.ImageLabels, &out.ImageLabels
			*out = make([]ImageLabel, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.ImageLabels = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildPostCommitSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildPostCommitSpec)
		out := out.(*BuildPostCommitSpec)
		if in.Command != nil {
			in, out := &in.Command, &out.Command
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Command = nil
		}
		if in.Args != nil {
			in, out := &in.Args, &out.Args
			*out = make([]string, len(*in))
			copy(*out, *in)
		} else {
			out.Args = nil
		}
		out.Script = in.Script
		return nil
	}
}

func DeepCopy_v1_BuildRequest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildRequest)
		out := out.(*BuildRequest)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_v1_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Revision = nil
		}
		if in.TriggeredByImage != nil {
			in, out := &in.TriggeredByImage, &out.TriggeredByImage
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.TriggeredByImage = nil
		}
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.From = nil
		}
		if in.Binary != nil {
			in, out := &in.Binary, &out.Binary
			*out = new(BinaryBuildSource)
			**out = **in
		} else {
			out.Binary = nil
		}
		if in.LastVersion != nil {
			in, out := &in.LastVersion, &out.LastVersion
			*out = new(int64)
			**out = **in
		} else {
			out.LastVersion = nil
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Env = nil
		}
		if in.TriggeredBy != nil {
			in, out := &in.TriggeredBy, &out.TriggeredBy
			*out = make([]BuildTriggerCause, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_BuildTriggerCause(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.TriggeredBy = nil
		}
		if in.DockerStrategyOptions != nil {
			in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
			*out = new(DockerStrategyOptions)
			if err := DeepCopy_v1_DockerStrategyOptions(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.DockerStrategyOptions = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildSource)
		out := out.(*BuildSource)
		out.Type = in.Type
		if in.Binary != nil {
			in, out := &in.Binary, &out.Binary
			*out = new(BinaryBuildSource)
			**out = **in
		} else {
			out.Binary = nil
		}
		if in.Dockerfile != nil {
			in, out := &in.Dockerfile, &out.Dockerfile
			*out = new(string)
			**out = **in
		} else {
			out.Dockerfile = nil
		}
		if in.Git != nil {
			in, out := &in.Git, &out.Git
			*out = new(GitBuildSource)
			if err := DeepCopy_v1_GitBuildSource(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Git = nil
		}
		if in.Images != nil {
			in, out := &in.Images, &out.Images
			*out = make([]ImageSource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ImageSource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Images = nil
		}
		out.ContextDir = in.ContextDir
		if in.SourceSecret != nil {
			in, out := &in.SourceSecret, &out.SourceSecret
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.SourceSecret = nil
		}
		if in.Secrets != nil {
			in, out := &in.Secrets, &out.Secrets
			*out = make([]SecretBuildSource, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Secrets = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildSpec)
		out := out.(*BuildSpec)
		if err := DeepCopy_v1_CommonSpec(&in.CommonSpec, &out.CommonSpec, c); err != nil {
			return err
		}
		if in.TriggeredBy != nil {
			in, out := &in.TriggeredBy, &out.TriggeredBy
			*out = make([]BuildTriggerCause, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_BuildTriggerCause(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.TriggeredBy = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStatus)
		out := out.(*BuildStatus)
		out.Phase = in.Phase
		out.Cancelled = in.Cancelled
		out.Reason = in.Reason
		out.Message = in.Message
		if in.StartTimestamp != nil {
			in, out := &in.StartTimestamp, &out.StartTimestamp
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.StartTimestamp = nil
		}
		if in.CompletionTimestamp != nil {
			in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.CompletionTimestamp = nil
		}
		out.Duration = in.Duration
		out.OutputDockerImageReference = in.OutputDockerImageReference
		if in.Config != nil {
			in, out := &in.Config, &out.Config
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.Config = nil
		}
		if err := DeepCopy_v1_BuildStatusOutput(&in.Output, &out.Output, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_BuildStatusOutput(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStatusOutput)
		out := out.(*BuildStatusOutput)
		if in.To != nil {
			in, out := &in.To, &out.To
			*out = new(BuildStatusOutputTo)
			**out = **in
		} else {
			out.To = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildStatusOutputTo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStatusOutputTo)
		out := out.(*BuildStatusOutputTo)
		out.ImageDigest = in.ImageDigest
		return nil
	}
}

func DeepCopy_v1_BuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStrategy)
		out := out.(*BuildStrategy)
		out.Type = in.Type
		if in.DockerStrategy != nil {
			in, out := &in.DockerStrategy, &out.DockerStrategy
			*out = new(DockerBuildStrategy)
			if err := DeepCopy_v1_DockerBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.DockerStrategy = nil
		}
		if in.SourceStrategy != nil {
			in, out := &in.SourceStrategy, &out.SourceStrategy
			*out = new(SourceBuildStrategy)
			if err := DeepCopy_v1_SourceBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.SourceStrategy = nil
		}
		if in.CustomStrategy != nil {
			in, out := &in.CustomStrategy, &out.CustomStrategy
			*out = new(CustomBuildStrategy)
			if err := DeepCopy_v1_CustomBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.CustomStrategy = nil
		}
		if in.JenkinsPipelineStrategy != nil {
			in, out := &in.JenkinsPipelineStrategy, &out.JenkinsPipelineStrategy
			*out = new(JenkinsPipelineBuildStrategy)
			if err := DeepCopy_v1_JenkinsPipelineBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.JenkinsPipelineStrategy = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildTriggerCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildTriggerCause)
		out := out.(*BuildTriggerCause)
		out.Message = in.Message
		if in.GenericWebHook != nil {
			in, out := &in.GenericWebHook, &out.GenericWebHook
			*out = new(GenericWebHookCause)
			if err := DeepCopy_v1_GenericWebHookCause(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.GenericWebHook = nil
		}
		if in.GitHubWebHook != nil {
			in, out := &in.GitHubWebHook, &out.GitHubWebHook
			*out = new(GitHubWebHookCause)
			if err := DeepCopy_v1_GitHubWebHookCause(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.GitHubWebHook = nil
		}
		if in.ImageChangeBuild != nil {
			in, out := &in.ImageChangeBuild, &out.ImageChangeBuild
			*out = new(ImageChangeCause)
			if err := DeepCopy_v1_ImageChangeCause(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.ImageChangeBuild = nil
		}
		return nil
	}
}

func DeepCopy_v1_BuildTriggerPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildTriggerPolicy)
		out := out.(*BuildTriggerPolicy)
		out.Type = in.Type
		if in.GitHubWebHook != nil {
			in, out := &in.GitHubWebHook, &out.GitHubWebHook
			*out = new(WebHookTrigger)
			**out = **in
		} else {
			out.GitHubWebHook = nil
		}
		if in.GenericWebHook != nil {
			in, out := &in.GenericWebHook, &out.GenericWebHook
			*out = new(WebHookTrigger)
			**out = **in
		} else {
			out.GenericWebHook = nil
		}
		if in.ImageChange != nil {
			in, out := &in.ImageChange, &out.ImageChange
			*out = new(ImageChangeTrigger)
			if err := DeepCopy_v1_ImageChangeTrigger(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.ImageChange = nil
		}
		return nil
	}
}

func DeepCopy_v1_CommonSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CommonSpec)
		out := out.(*CommonSpec)
		out.ServiceAccount = in.ServiceAccount
		if err := DeepCopy_v1_BuildSource(&in.Source, &out.Source, c); err != nil {
			return err
		}
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_v1_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Revision = nil
		}
		if err := DeepCopy_v1_BuildStrategy(&in.Strategy, &out.Strategy, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_BuildOutput(&in.Output, &out.Output, c); err != nil {
			return err
		}
		if err := api_v1.DeepCopy_v1_ResourceRequirements(&in.Resources, &out.Resources, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_BuildPostCommitSpec(&in.PostCommit, &out.PostCommit, c); err != nil {
			return err
		}
		if in.CompletionDeadlineSeconds != nil {
			in, out := &in.CompletionDeadlineSeconds, &out.CompletionDeadlineSeconds
			*out = new(int64)
			**out = **in
		} else {
			out.CompletionDeadlineSeconds = nil
		}
		if in.NodeSelector != nil {
			in, out := &in.NodeSelector, &out.NodeSelector
			*out = make(OptionalNodeSelector)
			for key, val := range *in {
				(*out)[key] = val
			}
		} else {
			out.NodeSelector = nil
		}
		return nil
	}
}

func DeepCopy_v1_CustomBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomBuildStrategy)
		out := out.(*CustomBuildStrategy)
		out.From = in.From
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.PullSecret = nil
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Env = nil
		}
		out.ExposeDockerSocket = in.ExposeDockerSocket
		out.ForcePull = in.ForcePull
		if in.Secrets != nil {
			in, out := &in.Secrets, &out.Secrets
			*out = make([]SecretSpec, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Secrets = nil
		}
		out.BuildAPIVersion = in.BuildAPIVersion
		return nil
	}
}

func DeepCopy_v1_DockerBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerBuildStrategy)
		out := out.(*DockerBuildStrategy)
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.From = nil
		}
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.PullSecret = nil
		}
		out.NoCache = in.NoCache
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Env = nil
		}
		out.ForcePull = in.ForcePull
		out.DockerfilePath = in.DockerfilePath
		if in.BuildArgs != nil {
			in, out := &in.BuildArgs, &out.BuildArgs
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.BuildArgs = nil
		}
		if in.ImageOptimizationPolicy != nil {
			in, out := &in.ImageOptimizationPolicy, &out.ImageOptimizationPolicy
			*out = new(ImageOptimizationPolicy)
			**out = **in
		} else {
			out.ImageOptimizationPolicy = nil
		}
		return nil
	}
}

func DeepCopy_v1_DockerStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerStrategyOptions)
		out := out.(*DockerStrategyOptions)
		if in.BuildArgs != nil {
			in, out := &in.BuildArgs, &out.BuildArgs
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.BuildArgs = nil
		}
		return nil
	}
}

func DeepCopy_v1_GenericWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GenericWebHookCause)
		out := out.(*GenericWebHookCause)
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_v1_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Revision = nil
		}
		out.Secret = in.Secret
		return nil
	}
}

func DeepCopy_v1_GenericWebHookEvent(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GenericWebHookEvent)
		out := out.(*GenericWebHookEvent)
		out.Type = in.Type
		if in.Git != nil {
			in, out := &in.Git, &out.Git
			*out = new(GitInfo)
			if err := DeepCopy_v1_GitInfo(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Git = nil
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Env = nil
		}
		if in.DockerStrategyOptions != nil {
			in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
			*out = new(DockerStrategyOptions)
			if err := DeepCopy_v1_DockerStrategyOptions(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.DockerStrategyOptions = nil
		}
		return nil
	}
}

func DeepCopy_v1_GitBuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitBuildSource)
		out := out.(*GitBuildSource)
		out.URI = in.URI
		out.Ref = in.Ref
		if err := DeepCopy_v1_ProxyConfig(&in.ProxyConfig, &out.ProxyConfig, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_GitHubWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitHubWebHookCause)
		out := out.(*GitHubWebHookCause)
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_v1_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		} else {
			out.Revision = nil
		}
		out.Secret = in.Secret
		return nil
	}
}

func DeepCopy_v1_GitInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitInfo)
		out := out.(*GitInfo)
		if err := DeepCopy_v1_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, c); err != nil {
			return err
		}
		out.GitSourceRevision = in.GitSourceRevision
		return nil
	}
}

func DeepCopy_v1_GitSourceRevision(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitSourceRevision)
		out := out.(*GitSourceRevision)
		out.Commit = in.Commit
		out.Author = in.Author
		out.Committer = in.Committer
		out.Message = in.Message
		return nil
	}
}

func DeepCopy_v1_ImageChangeCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageChangeCause)
		out := out.(*ImageChangeCause)
		out.ImageID = in.ImageID
		if in.FromRef != nil {
			in, out := &in.FromRef, &out.FromRef
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.FromRef = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageChangeTrigger(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageChangeTrigger)
		out := out.(*ImageChangeTrigger)
		out.LastTriggeredImageID = in.LastTriggeredImageID
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.From = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageLabel(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageLabel)
		out := out.(*ImageLabel)
		out.Name = in.Name
		out.Value = in.Value
		return nil
	}
}

func DeepCopy_v1_ImageSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageSource)
		out := out.(*ImageSource)
		out.From = in.From
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]ImageSourcePath, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.Paths = nil
		}
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.PullSecret = nil
		}
		return nil
	}
}

func DeepCopy_v1_ImageSourcePath(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageSourcePath)
		out := out.(*ImageSourcePath)
		out.SourcePath = in.SourcePath
		out.DestinationDir = in.DestinationDir
		return nil
	}
}

func DeepCopy_v1_JenkinsPipelineBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JenkinsPipelineBuildStrategy)
		out := out.(*JenkinsPipelineBuildStrategy)
		out.JenkinsfilePath = in.JenkinsfilePath
		out.Jenkinsfile = in.Jenkinsfile
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Env = nil
		}
		return nil
	}
}

func DeepCopy_v1_ProxyConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProxyConfig)
		out := out.(*ProxyConfig)
		if in.HTTPProxy != nil {
			in, out := &in.HTTPProxy, &out.HTTPProxy
			*out = new(string)
			**out = **in
		} else {
			out.HTTPProxy = nil
		}
		if in.HTTPSProxy != nil {
			in, out := &in.HTTPSProxy, &out.HTTPSProxy
			*out = new(string)
			**out = **in
		} else {
			out.HTTPSProxy = nil
		}
		if in.NoProxy != nil {
			in, out := &in.NoProxy, &out.NoProxy
			*out = new(string)
			**out = **in
		} else {
			out.NoProxy = nil
		}
		return nil
	}
}

func DeepCopy_v1_SecretBuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecretBuildSource)
		out := out.(*SecretBuildSource)
		out.Secret = in.Secret
		out.DestinationDir = in.DestinationDir
		return nil
	}
}

func DeepCopy_v1_SecretSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecretSpec)
		out := out.(*SecretSpec)
		out.SecretSource = in.SecretSource
		out.MountPath = in.MountPath
		return nil
	}
}

func DeepCopy_v1_SourceBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceBuildStrategy)
		out := out.(*SourceBuildStrategy)
		out.From = in.From
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api_v1.LocalObjectReference)
			**out = **in
		} else {
			out.PullSecret = nil
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api_v1.EnvVar, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Env = nil
		}
		out.Scripts = in.Scripts
		if in.Incremental != nil {
			in, out := &in.Incremental, &out.Incremental
			*out = new(bool)
			**out = **in
		} else {
			out.Incremental = nil
		}
		out.ForcePull = in.ForcePull
		if in.RuntimeImage != nil {
			in, out := &in.RuntimeImage, &out.RuntimeImage
			*out = new(api_v1.ObjectReference)
			**out = **in
		} else {
			out.RuntimeImage = nil
		}
		if in.RuntimeArtifacts != nil {
			in, out := &in.RuntimeArtifacts, &out.RuntimeArtifacts
			*out = make([]ImageSourcePath, len(*in))
			for i := range *in {
				(*out)[i] = (*in)[i]
			}
		} else {
			out.RuntimeArtifacts = nil
		}
		return nil
	}
}

func DeepCopy_v1_SourceControlUser(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceControlUser)
		out := out.(*SourceControlUser)
		out.Name = in.Name
		out.Email = in.Email
		return nil
	}
}

func DeepCopy_v1_SourceRevision(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceRevision)
		out := out.(*SourceRevision)
		out.Type = in.Type
		if in.Git != nil {
			in, out := &in.Git, &out.Git
			*out = new(GitSourceRevision)
			**out = **in
		} else {
			out.Git = nil
		}
		return nil
	}
}

func DeepCopy_v1_WebHookTrigger(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*WebHookTrigger)
		out := out.(*WebHookTrigger)
		out.Secret = in.Secret
		out.AllowEnv = in.AllowEnv
		return nil
	}
}
