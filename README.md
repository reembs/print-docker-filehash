### Docker file hash utility

Utility to help analyze docker cache invalidation issues. Run from your Dockerfile root and give the app a relative path, the output will include the hash as will be calculated by docker and the relevant file headers included.

Example usage:

```
bash-4.3# ./print-docker-filehash README.md
Hash: ddb18ace9975e1f5d541db0fadd7880da05e963d6926e919f46dbb2003d22976, Hdrs: [[name README.md] [mode 33188] [uid 501] [gid 20] [size 851] [typeflag 0] [linkname ] [uname ] [gname ] [devmajor 0] [devminor 0]]
```
