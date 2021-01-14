# apiserver

## 安装 apiserver-builder

目前 master 分支不能运行，应该安装 [v.18.0](https://github.com/kubernetes-sigs/apiserver-builder-alpha/releases/tag/v1.18.0)

## 实验步骤

1. 生成项目
``` shell
apiserver-boot init repo --domain core.kubrid.io
```

2. 创建 vmi resource
``` shell
 apiserver-boot create group version resource --group subresource --version v1alpha1 --non-namespaced=true --kind VirtualMachineInstance
```

3. 创建 poweraction subresource
``` shell
 apiserver-boot create subresource --subresource PowerAction --group subresource --version v1alpha1 --kind VirtualMachineInstance
 ```

4. 生成相关代码
``` shell
apiserver-boot build generated
```

5. 修改 cmd/manager/main.go 注释掉不需要的逻辑

6. 本地运行
``` shell
apiserver-boot run local --generate=false
```

7. 验证
``` shell
curl -k https://localhost:9443/apis/subresource.core.kubrid.io/v1alpha1/virtualmachineinstances
```

其生成的是一个完整的 apiserver，resource 存储在一个单独的 etcd 中。本实验中 apiserver-builder(v1.18.0) 自带了 etcd，并在本地运行了一个实例。这一点和我们目前的需求是不同的，我们是将 VMI 作为一个 CRD，通过一个独立的 apiserver 只是提供一些额外的操作。但是此处不得不先生成 VMI resource，否则无法添加 poweraction 等子资源。

VMI 的电源操作和 VNC 更接近 [pod/exec](https://github.com/kubernetes-sigs/apiserver-builder-alpha/blob/master/example/podexec/pkg/apis/podexec/exec_pod_rest.go) 这个示例 apiserver。pod/exec 实现了原生的 k8s apiserver 的 pod exec 功能。其实现上基本和 kubrid 之前的实现一样。定义了 `PodExecREST`，然后实现了 `Connecter()` 接口。

## apiserver-runtime

上面的实验是基于 v1.18.0 的 apiserver-builder。apiserver-builder master 的代码有较大的变化，其使用 [apiserver-runtime](https://github.com/kubernetes-sigs/apiserver-runtime/)。apiserver-runtime 是 kubebuilder 的一个子项目，现在处于很早期的阶段（不建议使用）。对 [kubernetes/apiserver](https://github.com/kubernetes/apiserver) 进行了一层封装，简化使用难度。可以类比 [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime/)

尝试使用 apiserver-runtime 来完成电源操作的逻辑没有成功：启动时其要求必须传入 etcd server 信息，遂放弃。后续应关注此项目，待其进一步成熟。

## 总结

目前还是推荐使用 kubernetes/apiserver 来构建 apiserver，从之前的代码来看，使用了 poweraction 子资源的方式后，代码量也不大。实现上可以参考 [pod/exec](https://github.com/kubernetes-sigs/apiserver-builder-alpha/blob/master/example/podexec/pkg/apis/podexec/exec_pod_rest.go) 例子，比较符合我们目前的需求。既能处理电源操作，也能处理 vnc。
