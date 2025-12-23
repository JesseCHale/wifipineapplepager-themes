# Hak5 Pineapple Pager Graphic Converter

Converts a PNG (480x222 pixels) to a pager graphic native framebuffer image.  The source image should be in "human" orientation - 480x222; this tool will automatically rotate it for the native framebuffer orientation on the pager.

The utility of this is somewhat limited since the native Pager UI will fight for the framebuffer, and Pager themes use standard graphics formats like PNG, however for those experimenting with making full screen graphics, this might save some time.

## Not for themes!

**This tool is not for making graphics for themes**!  This converts a full-screen PNG to a full-screen native framebuffer!

## Building

Install `go`, via packages, `brew`, etc

Run `make`

## Usage

`./pagergraphic [path to png file] [path to framebuffer file]`

The PNG will be converted to an appropriately rotated rgb565 framebuffer
