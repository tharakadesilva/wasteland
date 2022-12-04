load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "day1",
    srcs = ["solution.go"],
    importpath = "github.com/tharakadesilva/public/code-challenges/advent-of-code/2022/day1",
    visibility = ["//visibility:public"],
)

go_test(
    name = "day1_test",
    size = "small",
    srcs = ["solution_test.go"],
    data = glob(["testdata/**"]),
    embed = [":day1"],
)
