module hayashi/module2

go 1.18

replace hayashi/module1 => ../module1

require (
	hayashi/module1 v0.0.0-00010101000000-000000000000
	hayashi/module3/package2_module3 v0.0.0-00010101000000-000000000000
	hayashi/module3/package_module3 v0.0.0-00010101000000-000000000000
)

replace hayashi/module3/package_module3 => ../module3

replace hayashi/module3/package2_module3 => ../module3/sub
