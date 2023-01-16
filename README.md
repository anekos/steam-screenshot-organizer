# Steam Screenshot Organizer

Copy local screenshot files to specified directory including game ID.

![Icon](./icon.png)

# Usage

```
$ steam-screenshot-organizer.exe "N:\MyNAS\Screenshots"
```

... creates directories and files like below

```
$ tree "N:\MyNAS\Screenshots"
.
├── 1332010
│   ├── 1332010_20220722203537_1.png
│   ├── 1332010_20220722203556_1.png
│   ├── 1332010_20220722203632_1.png
│   └── 1332010_20220730210309_1.png
├── 1551360
│   ├── 20221230134954_1.jpg
│   └── 20230115205353_1.jpg
├── 292030
│   ├── 2018070800155498.png
│   ├── 20180923150200_1.png
│   ├── 20180923152422_1.png
│   ├── 20180923154041_1.png
│   ├── 20180923154044_1.png
...
```

You can include any title after the Game ID in the directory name and the command will recognize it.

```
.
├── 1332010 - Stray
│   ├── 1332010_20220722203537_1.png
│   ├── 1332010_20220722203556_1.png
│   ├── 1332010_20220722203632_1.png
│   └── 1332010_20220730210309_1.png
├── 1551360 - Forza Horizon 5
│   ├── 20221230134954_1.jpg
│   └── 20230115205353_1.jpg
├── 292030 - Witcher 3
│   ├── 2018070800155498.png
│   ├── 20180923150200_1.png
│   ├── 20180923152422_1.png
│   ├── 20180923154041_1.png
│   ├── 20180923154044_1.png
...
```