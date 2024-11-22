// Code generated by "core generate"; DO NOT EDIT.

package config

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cmd/core/config.Config", IDName: "config", Doc: "Config is the main config struct that contains all of the configuration\noptions for the Cogent Core command line tool.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Name", Doc: "Name is the user-friendly name of the project.\nThe default is based on the current directory name."}, {Name: "NamePrefix", Doc: "NamePrefix is the prefix to add to the default name of the project\nand any projects nested below it. A separating space is automatically included."}, {Name: "ID", Doc: "ID is the bundle / package ID to use for the project\n(required for building for mobile platforms and packaging\nfor desktop platforms). It is typically in the format com.org.app\n(eg: com.cogent.mail). It defaults to com.parentDirectory.currentDirectory."}, {Name: "About", Doc: "About is the about information for the project, which can be viewed via\nthe \"About\" button in the app bar. It is also used when packaging the app."}, {Name: "Version", Doc: "the version of the project to release"}, {Name: "Content", Doc: "Content, if specified, indicates that the app has core content pages\nlocated at this directory. If so, a directory tree will be made for all\nof the pages when building for platform web. This defaults to \"content\"\nwhen building an app for platform web that imports content."}, {Name: "Build", Doc: "the configuration options for the build, install, run, and pack commands"}, {Name: "Pack", Doc: "the configuration information for the pack command"}, {Name: "Web", Doc: "the configuration information for web"}, {Name: "Log", Doc: "the configuration options for the log and run commands"}, {Name: "Generate", Doc: "the configuration options for the generate command"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cmd/core/config.Build", IDName: "build", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Target", Doc: "the target platforms to build executables for"}, {Name: "Dir", Doc: "Dir is the directory to build the app from.\nIt defaults to the current directory."}, {Name: "Output", Doc: "Output is the directory to output the built app to.\nIt defaults to the current directory for desktop executables\nand \"bin/{platform}\" for all other platforms and command \"pack\"."}, {Name: "Debug", Doc: "whether to build/run the app in debug mode, which sets\nthe \"debug\" tag when building. On iOS and Android, this\nalso prints the program output."}, {Name: "IOSVersion", Doc: "the minimum version of the iOS SDK to compile against"}, {Name: "AndroidMinSDK", Doc: "the minimum supported Android SDK (uses-sdk/android:minSdkVersion in AndroidManifest.xml)"}, {Name: "AndroidTargetSDK", Doc: "the target Android SDK version (uses-sdk/android:targetSdkVersion in AndroidManifest.xml)"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cmd/core/config.Pack", IDName: "pack", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "DMG", Doc: "whether to build a .dmg file on macOS in addition to a .app file.\nThis is automatically disabled for the install command."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cmd/core/config.Log", IDName: "log", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Target", Doc: "the target platform to view the logs for (ios or android)"}, {Name: "Keep", Doc: "whether to keep the previous log messages or clear them"}, {Name: "All", Doc: "messages not generated from your app equal to or above this log level will be shown"}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cmd/core/config.Generate", IDName: "generate", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Enumgen", Doc: "the enum generation configuration options passed to enumgen"}, {Name: "Typegen", Doc: "the generation configuration options passed to typegen"}, {Name: "Dir", Doc: "the source directory to run generate on (can be multiple through ./...)"}, {Name: "Icons", Doc: "Icons, if specified, indicates to generate an icongen.go file with\nicon variables for the icon svg files in the specified folder."}}})

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/cmd/core/config.Web", IDName: "web", Doc: "Web containts the configuration information for building for web and creating\nthe HTML page that loads a Go wasm app and its resources.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Fields: []types.Field{{Name: "Port", Doc: "Port is the port to serve the page at when using the serve command."}, {Name: "RandomVersion", Doc: "RandomVersion is whether to make the app worker version random.\nIt is enabled by default and should be kept on for easy deployment."}, {Name: "Gzip", Doc: "Gzip is whether to gzip the app.wasm file that is built in the build command\nand serve it as a gzip-encoded file in the run command."}, {Name: "GenerateHTML", Doc: "GenerateHTML is whether to generate an HTML version of app content for\npreview and SEO purposes."}, {Name: "Lang", Doc: "The page language.\n\nDEFAULT: en."}, {Name: "Author", Doc: "The page authors."}, {Name: "Keywords", Doc: "The page keywords."}, {Name: "Image", Doc: "The path of the default image that is used by social networks when\nlinking the app."}, {Name: "AutoUpdateInterval", Doc: "The interval between each app auto-update while running in a web browser.\nZero or negative values deactivates the auto-update mechanism.\n\nDefault is 10 seconds."}, {Name: "VanityURL", Doc: "If specified, make this page a Go import vanity URL with this\nmodule URL, pointing to the GitHub repository specified by GithubVanityURL\n(eg: cogentcore.org/core)."}, {Name: "GithubVanityRepository", Doc: "If VanityURL is specified, the underlying GitHub repository for the vanity URL\n(eg: cogentcore/core)."}, {Name: "Env", Doc: "The environment variables that are passed to the progressive web app.\n\nReserved keys:\n- GOAPP_STATIC_RESOURCES_URL"}, {Name: "WasmContentLengthHeader", Doc: "The HTTP header to retrieve the WebAssembly file content length.\n\nContent length finding falls back to the Content-Length HTTP header when\nno content length is found with the defined header."}}})
