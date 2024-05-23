type DatePattern = 'YYYY' | 'MM' | 'DD' | 'HH' | 'mm' | 'ss' | 'M' | 'D' | 'H' | 'm' | 's';
export const format = function (date: Date, pattern = 'YYYY-MM-DD HH:mm:ss') {
  const year = date.getFullYear(),
    M = date.getMonth() + 1,
    D = date.getDate(),
    H = date.getHours(),
    m = date.getMinutes(),
    s = date.getSeconds();
  const data = {
    YYYY: year,
    MM: `${M}`.padStart(2, '0'),
    DD: `${D}`.padStart(2, '0'),
    HH: `${H}`.padStart(2, '0'),
    mm: `${m}`.padStart(2, '0'),
    ss: `${s}`.padStart(2, '0'),
    M,
    D,
    H,
    m,
    s,
  };
  for (let key in data) {
    pattern = pattern.replace(new RegExp(key), data[key as DatePattern] as string);
  }
  return pattern;
};
Date.prototype.format = function (pattern = 'YYYY-MM-DD HH:mm:ss') {
  return format(this, pattern);
};
