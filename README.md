# ms
Ported from https://github.com/zeit/ms

[![Build Status](https://travis-ci.org/fishead/ms.svg?branch=master)](https://travis-ci.org/fishead/ms)

Use this package to easily convert various time formats to milliseconds.

## Examples

```js
Parse('2 days')  // 172800000
Parse('1d')      // 86400000
Parse('10h')     // 36000000
Parse('2.5 hrs') // 9000000
Parse('2h')      // 7200000
Parse('1m')      // 60000
Parse('5s')      // 5000
Parse('1y')      // 31557600000
Parse('100')     // 100
Parse('-3 days') // -259200000
Parse('-1h')     // -3600000
Parse('-200')    // -200
```

### Convert from Milliseconds

```js
FmtShort(60000)             // "1m"
FmtShort(2 * 60000)         // "2m"
FmtShort(-3 * 60000)        // "-3m"
FmtShort(Parse('10 hours')) // "10h"
```

### Time Format Written-Out

```js
FmtLong(60000)              // "1 minute"
FmtLong(2 * 60000)          // "2 minutes"
FmtLong(-3 * 60000)         // "-3 minutes"
FmtLong(Parse('10 hours'))  // "10 hours"
```

## Features

- If a number is supplied to `ms`, a string with a unit is returned
- If a string that contains the number is supplied, it returns it as a number (e.g.: it returns `100` for `'100'`)
- If you pass a string with a number and a valid unit, the number of equivalent milliseconds is returned

## Thanks
- [ms](https://github.com/zeit/ms)
