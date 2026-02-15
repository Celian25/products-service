export const formatCurrency = (
  amount: number,
  currency: string = "RUB",
  locale: string = "ru-RU",
): string => {
  return new Intl.NumberFormat(locale, {
    style: "currency",
    currency: currency,
    minimumFractionDigits: 0,
    maximumFractionDigits: 2,
    currencyDisplay: "narrowSymbol", // 'symbol', 'narrowSymbol', 'code', 'name'
  }).format(amount);
};
