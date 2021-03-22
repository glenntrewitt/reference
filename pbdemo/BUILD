load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "pbdemo_proto",
    srcs = ["pbdemo.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "pbdemo_go_proto",
    importpath = "github.com/glenntrewitt/reference/pbdemo",
    proto = ":pbdemo_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "pbdemo",
    embed = [":pbdemo_go_proto"],
    importpath = "github.com/glenntrewitt/reference/pbdemo",
    visibility = ["//visibility:public"],
)
