#ifndef LIST_H
#define LIST_H
#include "date.h"

typedef struct Node Node;
struct Node {
    Date* d;
    Node* prev;
    Node* next;
};

typedef struct List List;
struct List {
    int size;
    Node* head;
    Node* tail;
};

List* newList();
void push_back(List* l, Date* d);
void push_front(List* l, Date* d);
void insert_after(List* l, int pos, Date* d);
void insert_before(List* l, int pos, Date* d);
Node* search(List* l, int pos);
void erase(List* l, int pos);
void printList(List* l);
void freeList(List* l);

#endif