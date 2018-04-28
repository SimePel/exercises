#ifndef DATE_H
#define DATE_H

struct Date {
    int day;
    int month;
    int year;
};
typedef struct Date Date;

enum Str_rep {
    Digital,    // 30.01.2001
    Brief,      // 30 Янв 01
    Full       // 30 января 2001 г.
};
typedef enum Str_rep Str_rep;

void print(const Date* d, const Str_rep rep);
void add_day(Date *d);

#endif // DATE_H