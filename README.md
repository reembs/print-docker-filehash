### Docker file hash utility

Utility to help analyze docker cache invalidation issues. Run from your Dockerfile root and give the app a relative path, the output will include the hash as will be calculated by docker and the relevant file headers included.

Example usage:

```
bash-4.3# ./print-docker-filehash create-jar.sh
Path (relative to dockerfile): create-jar.sh
Hash: 80097707e89b5f454d9cd768115a8dccc9f76c870de129856e6345b2a260f524
Hdrs: [[name create-jar.sh] [mode 33277] [uid 0] [gid 0] [size 286] [typeflag 0] [linkname ] [uname ] [gname ] [devmajor 0] [devminor 0]]
```

Included Dockerfile will compile a native Ubuntu executable to run anywhere. Run with following command. After execution the binary will be available under ./result directory.

```
docker build -t gocompile .
docker run -v ${PWD}/result:/go/result -ti gocompile
```
