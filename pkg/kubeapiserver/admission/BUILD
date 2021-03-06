package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = [
        "config_test.go",
        "init_test.go",
    ],
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//vendor:k8s.io/apiserver/pkg/admission",
        "//vendor:k8s.io/apiserver/pkg/apis/apiserver",
        "//vendor:k8s.io/apiserver/pkg/apis/apiserver/install",
        "//vendor:k8s.io/apiserver/pkg/authorization/authorizer",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "initializer.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/client/clientset_generated/internalclientset:go_default_library",
        "//pkg/controller/informers:go_default_library",
        "//vendor:github.com/ghodss/yaml",
        "//vendor:github.com/golang/glog",
        "//vendor:k8s.io/apimachinery/pkg/runtime",
        "//vendor:k8s.io/apimachinery/pkg/util/sets",
        "//vendor:k8s.io/apiserver/pkg/admission",
        "//vendor:k8s.io/apiserver/pkg/apis",
        "//vendor:k8s.io/apiserver/pkg/apis/apiserver",
        "//vendor:k8s.io/apiserver/pkg/apis/apiserver/install",
        "//vendor:k8s.io/apiserver/pkg/apis/apiserver/v1alpha1",
        "//vendor:k8s.io/apiserver/pkg/authorization/authorizer",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
