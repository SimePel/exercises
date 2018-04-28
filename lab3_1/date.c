#include "stdio.h"
#include "stdbool.h"
#include "date.h"

const int days_in_leap_year[] = {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
const int days_in_not_leap_year[] = {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
char* full_months[] = {"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"};
char* brief_months[] = {"Янв", "Фев", "Мрт", "Апр", "Май", "Июн", "Июл", "Авг", "Сен", "Окт", "Нбр", "Дек"};

static char* to_brief_str(const int m);
static char* to_full_str(const int m);
static void inc_month(Date *d);
static bool is_leap_year(const int year);

void print(const Date* d, const Str_rep rep) {
    switch(rep) {
    case Digital:
        printf("%d.%d.%d\n", d->day, d->month, d->year);
        break;
    case Brief:
        printf("%d.%s.%d", d->day, to_brief_str(d->month), d->year);
        break;
    case Full:
        printf("%d.%s.%d\n", d->day, to_full_str(d->month), d->year);
        break;
    default:
        printf("unrecognized representation type\n");
    }
}

static char* to_brief_str(const int m) {
    return brief_months[m-1];
}

static char* to_full_str(const int m) {
    return full_months[m-1];
}

// Добавляет один день и следит за корректностью
void add_day(Date *d) {
    d->day++;
    if (is_leap_year(d->year)) {    
        if (d->day > days_in_leap_year[d->month - 1]) {
            inc_month(d);
        }
    }
    else {
        if (d->day > days_in_not_leap_year[d->month - 1]) {
            inc_month(d);
        }
    }
}

static void inc_month(Date *d) {
    d->month++;
    if (d->month > 12) {
        d->year++;
        d->month = 1;
    }
    d->day = 1;
}

static bool is_leap_year(const int year) {
    return year % 4 == 0 && (year % 100 != 0 || year % 400 == 0);
}