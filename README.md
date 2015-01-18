## reddit.com/r/dailyprogrammer - Challenge #197 - Easy

http://www.reddit.com/r/dailyprogrammer/comments/2s7ezp/20150112_challenge_197_easy_isbn_validator/cns9v7c

I had a bit of fun with this. 

Implemented:
* ISBN-10 validation
* ISBN-10 check digit calculation
* ISBN-10 amazon lookup
* ISBN-10 neighbors (10) calculation & lookup

Example output:

```
    dpc197 $ source env_example.sh    

    dpc197 $ go run isbn.go -i 0761174427 -n
    Checking for valid ISBN: 0761174427
    ISBN-10: 0761174427 www.amazon.com/Unlikely-Loves-Heartwarming-Stories-Kingdom/dp/0761174427
    -----   Neighbor ISBNs  -----
    0761174400 www.amazon.com/Hero-Dogs-2014-Wall-Calendar/dp/0761174400
    0761174419 www.amazon.com/Unlikely-Heroes-Inspiring-Stories-Courage/dp/0761174419
    0761174427 www.amazon.com/Unlikely-Loves-Heartwarming-Stories-Kingdom/dp/0761174427
    0761174435
    0761174443
    0761174451
    076117446X
    0761174478
    0761174486 www.amazon.com/Moms-Family-2014-Desk-Planner/dp/0761174486
    0761174494 www.amazon.com/Lego-Calendar-2014-Workman-Publishing/dp/0761174494
    ----- ----- ----- ----- -----

    dpc197 $ go run isbn.go -i 076117442X -n
    Checking for valid ISBN: 076117442X
    Not Valid ISBN-10:  Invalid check digit: expected (7) received (X)
    Looking up expected ISBN-10: 0761174427
    ISBN-10: 0761174427 www.amazon.com/Unlikely-Loves-Heartwarming-Stories-Kingdom/dp/0761174427 
    -----   Neighbor ISBNs  -----
    0761174400 www.amazon.com/Hero-Dogs-2014-Wall-Calendar/dp/0761174400
    0761174419 www.amazon.com/Unlikely-Heroes-Inspiring-Stories-Courage/dp/0761174419
    0761174427 www.amazon.com/Unlikely-Loves-Heartwarming-Stories-Kingdom/dp/0761174427
    0761174435
    0761174443
    0761174451
    076117446X
    0761174478
    0761174486 www.amazon.com/Moms-Family-2014-Desk-Planner/dp/0761174486
    0761174494 www.amazon.com/Lego-Calendar-2014-Workman-Publishing/dp/0761174494
    ----- ----- ----- ----- -----
```
