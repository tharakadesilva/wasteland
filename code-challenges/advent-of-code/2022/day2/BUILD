load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "day2",
    srcs = ["solution.go"],
    importpath = "github.com/tharakadesilva/public/code-challenges/advent-of-code/2022/day2",
    visibility = ["//visibility:public"],
)

go_test(
    name = "day2_test",
    size = "small",
    srcs = ["solution_test.go"],
    data = [
        "inputs/input1.txt",
        "inputs/input2.txt",
    ],
    embed = [":day2"],
)
