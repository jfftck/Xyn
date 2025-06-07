# Xyn Programming Language

Xyn is a modern programming language designed for simplicity and efficiency, offering a streamlined syntax and innovative features. It is particularly suited for asynchronous programming and ensures code safety by eliminating null references. Below are the core features of Xyn and an explanation of the provided example.

> [!IMPORTANT]
> This is in early development and is currently unstable.

My hope is this language might inspire others to help in further developing Xyn, but this is really just a pet project of mine and I don't expect anything more for it.

---

## Key Features

### 1. **Signals** (Asynchronous Communication)
Xyn introduces a feature called **signals** for asynchronous programming. Signals allow values to be emitted to multiple or single inputs, enabling efficient communication between different parts of an application.

### 2. **No Nulls**
Xyn does not support null values. Instead, variables can be declared as "maybe" values using a question mark (`?`) at the end of the variable name. Accessing such variables requires an `else` expression to handle cases where the value is not set, ensuring safer code and avoiding null pointer exceptions. To set to a "none" state, you just need to assign the variable to a question mark (`?`).

Example:
```xyn
<string> page? // Declares a maybe value, it requires a type to be given
page? else "defaultPage" // Accesses the value with a fallback
page? = ? // Set the state of the maybe to unset, this will clear any value
```

### 3. **Errors as Values**
Xyn 

---

## Example Code

Below is an example of a server application written in Xyn:

```xyn
{ httpServer } <StatusCode, Get, Post, Put, Delete> := @import @std.http;
{ consoleLogger, fileLogger, multiLogger } <Info, Debug, Warn, Error, Panic, ConsoleLogger, FileLogger> := @import @std.logger;
{ argsParser, cpu } <Args, CORE_TYPE> := @import @std.os;

{ createPost, deletePost, getPost, getPosts, updatePost } := @import @p`./routes`;
{ handleAuth } := @import @p`./auth`;

@export type enum Status {
	Init,
	Starting,
	Running,
	Stopping,
	Stopped,
}

@export type struct AppUpdates {
	<Status> status,
}

public const DEFAULT_PORT = 8080;

@log(@@multiLogger.new<ConsoleLogger | FileLogger>) -> match! {
	(<ConsoleLogger>, <Info | Debug | Warn | Error | Panic>) ex => ex -> consoleLogger.newOrLog,
	(<FileLogger>, <Warn | Error | Panic>) ex => ex -> fileLogger.newOrLog
}

@main(@@cpu.useFreeCore(coreType = .eCore) as #mainCore) fn(<Args> args)
		[argsParser, httpServer, createPost, deletePost, getPost, getPosts, updatePost] {
	options? <:= argsParser(args) -> match! {
		(<str> “port” | “p”, <int> userPort) => :{ port: userPort }
	}

	<Status?> currentStatus?;
	@signal(<Status?> status?) fn!(<AppUpdates> payload) <Status> {
		newStatus := match (status?, payload.status) {
			(.Init, .Starting) => .Starting,
			(.Starting, .Running) => .Running,
			(.Running, .Stopping) => .Stopping,
			(.Stopping, .Stopped) => .Stopped,
			_ => status? else .Init
		}

		yeild newStatus where != payload.status;
	}
	@set(<Status?> status?) fn(<Status> newStatus) {
		status? <- fn![newStatus] <AppUpdates> {
			return @new<AppUpdate>() { status = newStatus };
		}
	}

	currentStatus? -> fn!(<Status> status) {
		print(<str(status)>);
	}

	currentStatus? = .Init;
	app := httpServer.createPool(cpu.threads() * cpu.cores() // 2, excludedCores=[#mainCore]);

	app.staticRoute("/", @p`./static`);

	api := app.group(@p`api/`);
	v1 := api.group(@p`v1/`);

	v1.route<Get | Post | Put | Delete>(@p`posts/<str?>`) <-> match! {
		(<Get>, '') => getPosts(),
		(<Get>, <str> id) => getPost(id),
		(<Post>, '') as req where && handleAuth(req) => createPost(req.body),
		(<Put>, <str> id) as req where && handleAuth(req) => updatePost(id, req.body),
		(<Delete>, <str> id) as req where && handleAuth(req) => deletePost(id),
		_ => app.httpError(“Invalid access to posts”),
	}

	currentStatus? = .Starting;
	app.run(options?.port else DEFAULT_PORT) -> fn!(<StatusCode> ex) { @log error(`Failed to load page: ${ex.message}`); }

	currentStatus? = .Running;

	while currentStatus? == .Running {

	}
}
```

---

## Explanation of the Example

### Import Statements

```xyn
{ httpServer } <StatusCode> := import @std.http;
{ consoleLogger, fileLogger, multiLogger } <Info, Debug, Warn, Error, Panic> := import @std.logger;
{ argsParser, cpu } <Args> := import @std.os;

{ createPost, deletePost, getPost, getPosts, updatePost } := import p`./routes`;
```

- Xyn uses structured imports to bring in modules and their specific features.
- For example, httpServer is imported from the @std.http module with the associated StatusCode type.
- Multiple loggers (consoleLogger, fileLogger, multiLogger) are imported with their logging level types.

### Enum and Struct Definitions

```xyn
export enum Status {
    Init,
    Starting,
    Running,
    Stopping,
    Stopped,
}

export struct AppUpdates {
    Status status,
}
```

- Status is an enumeration representing different application states.
- AppUpdates is a structure containing the Status field, used to update the application's state.

### Constants and Logging

```xyn
const DEFAULT_PORT = 8080;

@log = multiLogger.new -> match! {
    <Info | Debug | Warn | Error | Panic> ex => ex -> consoleLogger.newOrLog,
    <Warn | Error | Panic> ex => ex -> fileLogger.newOrLog
}
```

- DEFAULT_PORT is a constant for the default server port.
- The @log signal uses multiLogger to handle logging. It routes messages to consoleLogger for all levels and additionally to fileLogger for warnings, errors, and panics.

### Main Function

```xyn
@main = @@cpu.useFirstCore fn(<Args> args)
        [argsParser, httpServer, createPost, deletePost, getPost, getPosts, updatePost] {
    // Function logic here
}
```

- The @main function is the entry point of the application. It runs on the first CPU core and uses the provided modules (argsParser, httpServer, etc.).

### Parsing Arguments

```xyn
options? <:= argsParser(args) -> match! {
    port: (“port” | “p”) as userPort => userPort
}
```

- Command-line arguments are parsed using argsParser. The options? variable is a maybe value that holds the user-provided port or falls back to DEFAULT_PORT.

### Managing Application Status

```xyn
<Status> currentStatus?;
currentStatus?.@signal = fn!(<AppUpdates> payload) <Status> {
    // Status update logic
}
```

- currentStatus? is a maybe value representing the application's current status.
- A signal is defined to update the currentStatus based on incoming AppUpdates.

### Setting Up the HTTP Server

```xyn
app := httpServer.createPool(cpu.threads() * cpu.cores() / 2, excludedCores=[cpu.firstCore]);

app.staticRoute("/") <-> app.render(p`./static`);
```

- The HTTP server is created with a thread pool based on available CPU cores, excluding the first core.
- A static route is set up to serve files from the ./static directory.

### Defining API Routes

```xyn
api := app.group(p`api`);
v1 := api.group(p`v1`);

v1.router(p`posts/`) <-> match! {
    app.get() => getPosts(),
    app.get(p`${id}`).pre(auth) => getPost(id),
    app.post().pre(auth) as req => createPost(req.body),
    app.put(p`${id}`).pre(auth) as req => updatePost(id, req.body),
    app.delete(p`${id}`).pre(auth) => deletePost(id),
    _ => app.apiError(“Invalid access to posts”),
}
```

- API routes are grouped under /api/v1/posts/.
- Each route is mapped to a handler function (getPosts, getPost, etc.).
- The auth middleware is applied to routes requiring authentication.

### Running the Server

```xyn
for app.run(options?.port else DEFAULT_PORT) -> fn!(<StatusCode> ex) { 
    @log error(`Failed to load page: ${ex.message}`); 
} {
    currentStatus? -> fn!(<Status> status) {
        print(str(status));
    }
}
```

- The server runs on the user-specified port or the default port.
- Errors during server execution are logged using the @log signal.
- The currentStatus? signal prints the application's status whenever it changes.

---

## Conclusion

Xyn simplifies asynchronous programming with its signals feature and ensures code safety by eliminating null references. The provided example showcases how to build a robust server application using Xyn's features, including structured imports, maybe values, and efficient signal handling.
