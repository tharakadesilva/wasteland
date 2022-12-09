load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "day9",
    srcs = ["solution.go"],
    importpath = "github.com/tharakadesilva/public/code-challenges/advent-of-code/2022/day9",
    visibility = ["//visibility:public"],
)

go_test(
    name = "day9_test",
    size = "small",
    srcs = ["solution_test.go"],
    data = glob(["testdata/**"]),
    embed = [":day9"],
)
