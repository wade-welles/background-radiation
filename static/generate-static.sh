#!/bin/bash
for i in {1..24}
do
    mx=3840;my=2160;head -c "$((3*mx*my))" /dev/urandom | convert -depth 8 -size "${mx}x${my}" RGB:- random$i.png
done
