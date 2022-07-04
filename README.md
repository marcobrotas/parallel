# parallel
Simple linux tool to run commands in parallel

### Installation
 - run command `bash build`
 - move the resulting `parallel` binary to your `~/bin`

### Usage
```shell
parallel 'echo "foo"' 'echo "bar"'
```

### Notes
Keep in mind the tool is not prepared to handle output streams per command.  
It is solely intended to provide a simple and effective way to execute multiple commands in parallel (depending on your CPU).  
It will keep track and display the output of each command accordingly, and it will wait for all the commands to finish execution.  
Once done, it will return successfully only if all commands were successful.  
Otherwise, the result will be an error code and an indication of which commands resulted in an error.

example success:
```shell
YYYY/MM/DD hh:mm:ss Done!
```

example failure:
```shell
YYYY/MM/DD hh:mm:ss Execution finished but the following commands returned errors: [Command 2]
```
