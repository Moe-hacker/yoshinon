<p align="center">
    <img src="https://github.com/Moe-hacker/yoshinon/raw/main/logo.png", width="75%"/>
</p>

## About:       
Yoshinon ~~(よしのん)~~ is a whiptail-like dialog box written with Bubble Tea (Go).          
The command-line usage is very similar to dialog/whiptail.      
## Usage:
```
Box options:

  --msgbox      [options] <text> <height> <width>
  --yesno       [options] <text> <height> <width>
  --infobox     [options] <text> <height> <width>
  --inputbox    [options] <text> <height> <width> [init]
  --passwordbox [options] <text> <height> <width>
  --menu        [options] <text> <height> <width> <listheight> [tag item] ...
  --checklist   [options] <text> <height> <width> <listheight> [tag item status]...
  --radiolist   [options] <text> <height> <width> <listheight> [tag item status]...
  --gauge       [options] <text> <height> <width>

Options:

  --border  [rounded/normal/thick/double/hidden]
  --help    print this message
  --version print version information
```     
## TODO:
- [X] msgbox       
- [X] yesno  
- [X] inputbox
- [X] passwordbox
- [X] menu
- [X] checklist
- [X] radiolist
- [X] gauge
- [X] help
- [ ] man page

## Thanks:
[Bubble Tea](https://github.com/charmbracelet/bubbletea)      
```
●   ●  ●●●   ●●●  ●   ●  ●●●  ●   ●  ●●●  ●   ●
 ● ●  ●   ● ●     ●   ●   ●   ●●  ● ●   ● ●●  ●
  ●   ●   ●  ●●●  ●●●●●   ●   ● ● ● ●   ● ● ● ●
  ●   ●   ●     ● ●   ●   ●   ●  ●● ●   ● ●  ●●
  ●    ●●●   ●●●  ●   ●  ●●●  ●   ●  ●●●  ●   ●
```
<p align="center">「鸟になって云をつかんで、</p>    
<p align="center">风になって遥远くへ</p>   
<p align="center">希望を抱いて飞んだ」</p>   