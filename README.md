# Go Api Starter
 ## A base skeleton for go api<br><br>
 This project is intended to be used a base starting point to create rest api,
 the scripts included will result in a docker image/container exposing the 8080 port.
 <br><br>
 Some useful functions and libraries are included (log, env vars, etc...), i decided to not include any web framework/library, the choise is yours.

### Build

To build the project image, run the following command:

- Linux/Mac:
```sh
./build_scripts/build.sh
```
- Windows (wsl installed)
```sh
sh ./build_scripts/build.sh
```
### Run

To run the docker container:
- Linux/Mac:
```sh
./build_scripts/run.sh
```
- Windows (wsl installed)
```sh
sh ./build_scripts/run.sh
```
### Test coverage report

To get the test coverage report, running this command will generate a .out file in the root folder:
- Linux/Mac:
```sh
./build_scripts/test.sh
```
- Windows (wsl installed)
```sh
sh ./build_scripts/test.sh
```
