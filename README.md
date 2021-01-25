# pitop 
<div align="center">

**Raspberry Pi terminal based activity monitor**


<img src="./assets/pitop.gif" />


</div>

---

Yes I know there are plenty of solutions already available, but I wanted to develop my own terminal based activity monitor with Go, for my Raspberry Pi.

This is for RPI 4, it should work on RPI 3 (Update : It works on RPI 3). 


## Install 

**Note**: Prebuilt binaries for 32-bit and 64-bit operating systems, doesn't require Go.

### 32-bit

- Raspberry Pi OS 32-bit
- Raspian 
- RetropPie Raspberry Pi 4/400 

```bash 
curl -sSL https://raw.githubusercontent.com/PierreKieffer/pitop/master/install/install_pitop32.sh | bash
```
### 64-bit

- Ubuntu server 20.04 LTS 64-bit for Raspberry Pi

```bash 
curl -sSL https://raw.githubusercontent.com/PierreKieffer/pitop/master/install/install_pitop64.sh | bash
```

## Run 
```bash
pitop
```
## Built With

- [gizak/termui](https://github.com/gizak/termui)
  - [nsf/termbox](https://github.com/nsf/termbox-go)

No external package is used for system data extraction and manipulation. 

## License 
BSD



