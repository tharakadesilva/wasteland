load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "day4",
    srcs = ["solution.go"],
    importpath = "github.com/tharakadesilva/public/code-challenges/advent-of-code/2022/day4",
    visibility = ["//visibility:public"],
)

go_test(
    name = "day4_test",
    size = "small",
    srcs = ["solution_test.go"],
    data = glob(["testdata/**"]),
    embed = [":day4"],
)
