cc_test(
  name = "hello_test",
  size = "small",
  srcs = ["hello_test.cc"],
  deps = ["@com_google_googletest//:gtest_main"],
)

cc_library(
  name = "sebastien",
  srcs = ["sebastien.cc"],
  hdrs = ["include/sebastien.h"],
  deps = [":tank"],
)

cc_library(
  name = "tank",
  hdrs = ["include/tank.h"],
)

cc_library(
  name = "algorithm",
  srcs = ["algorithm.cc"],
  hdrs = ["include/algorithm.h"],
  deps = [
    ":tank",
    ":sebastien",
  ]
)
