# Deskapps

This application lists the currently installed desktop applications on a Linux
distribution. It searches the default directories where the desktop files are
usually placed (or uses the ones given in the command line), parses the files
in those directories and shows the results.

Custom desktop file directories can be given with the `-d` or `--deskDir`
command line switch. To add multiple entries, the switch also needs to be
repeated:

```bash
./deskapps -d /usr/share/applications -d ~/.local/share/applications
```

As an example of how it looks like, here's a short extract of an actual output:

```
[ConsoleOnly]
	ranger (ranger)
		Launches the ranger file manager
	Htop (htop)
		Show System Processes
	Vifm (vifm %F)
		Vi[m] like ncurses based file manager
	ToolBox (bmenu)
		Collection Of Terminal Applications In A Simple UI
	Vifm (/usr/bin/alacritty -e vifm %F)
		Vi[m] like ncurses based file manager

[Office]
	LibreOffice Draw (libreoffice --draw %U)
		Create and edit drawings, flow charts, and logos by using Draw.
	LibreOffice Math (libreoffice --math %U)
		Create and edit scientific formulas and equations by using Math.
	LibreOffice (libreoffice %U)
		The office productivity suite compatible to the open and standardized ODF document format. Supported by The Document Foundation.
	Foliate (com.github.johnfactotum.Foliate %U)
		View eBooks
	LibreOffice Base (libreoffice --base %U)
		Manage databases, create queries and reports to track and manage your information by using Base.
	LibreOffice Calc (libreoffice --calc %U)
		Perform calculations, analyze information and manage lists in spreadsheets by using Calc.
	Calibre (calibre --detach %U)
		E-book library management: Convert, view, share, catalogue all your e-books
	ePDFViewer (epdfview %f)
		Lightweight PDF document viewer
	E-book editor (ebook-edit --detach %f)
		Edit E-books in various formats
	Zathura (zathura %U)
		A minimalistic document viewer
	LibreOffice Writer (libreoffice --writer %U)
		Create and edit text and images in letters, reports, documents and Web pages by using Writer.
	LibreOffice Impress (libreoffice --impress %U)
		Create and edit presentations for slideshows, meeting and Web pages by using Impress.
	LRF viewer (lrfviewer %f)
		Viewer for LRF files (SONY ebook format files)
	E-book viewer (ebook-viewer --detach %f)
		Viewer for E-books in all the major formats
	Zathura (zathura %U)
		A minimalistic document viewer
```

## Compilation

The application needs to be compiled. For this, a working Go build environment
is needed. Please refer to your distribution's package manager or visit
[this][1] page for information on how to install Go.

>	**Note**
>
>	Go can typically be installed using package managers, although it might not
>	be the latest version that's available (especially if the package manager
>	serves a point release distribution). Some install commands for popular
>	package managers are:
>
>	Arch/Manjaro: `packman -S go`
>	Fedora: `dnf install golang`
>	Ubuntu: `apt-get install golang`

Once Go is available on the system, the compilation can be performed by
issuing:

```bash
go build -mod=mod -o build/deskapps ./cmd/deskapps
```

or by using the attached Makefile:

```bash
make all
```

The binary will be available in the `build` folder of the repositry root.

## Installation

The application can also be installed from `github` using the following
command:

```bash
go install github.com/nagygr/deskapps/cmd/deskapps@latest
```

>	**Note**
>
>	Please note, that the command above also requires Go to be installed on the
>	system. Please see details about Go installation above.

[1]: https://go.dev/doc/install
