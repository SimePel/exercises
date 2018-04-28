#include "stdio.h"
#include "stdlib.h"
#include "list.h"

static void addNodeBetween(Node* lhs, Node* rhs, Date* d);
static Node* newNode(Date* d);
static void deleteNode(Node* lhs, Node* n, Node* rhs);

// newList returns new list with allocated head and tail
List* newList() {
    List* l = malloc(sizeof(List));
    l->head = malloc(sizeof(Node));
    l->tail = malloc(sizeof(Node));

    l->head->prev = NULL;
    l->head->next = l->tail;
    l->tail->prev = l->head;
    l->tail->next = NULL;

    l->size = 0;
    return l;
}

// push_back adds Date before the tail of the List
void push_back(List* l, Date* d) {
    addNodeBetween(l->tail->prev, l->tail, d);
    l->size++;
}

// push_front adds Date after the head of the List
void push_front(List* l, Date* d) {
    addNodeBetween(l->head, l->head->next, d);
    l->size++;
}

// insert_after adds Date after pos. Doesn't work if List is empty,
// in that case use push_front or push_back
void insert_after(List* l, int pos, Date* d) {
    Node* n = search(l, pos);
    if (n == NULL) {
        printf("search returns NULL\n");
        return;
    }
    addNodeBetween(n, n->next, d);
    l->size++;
}

// insert_before adds Date after pos. Doesn't work if List is empty,
// in that case use push_front or push_back
void insert_before(List* l, int pos, Date* d) {
    Node* n = search(l, pos);
    if (n == NULL) {
        printf("search returns NULL\n");        
        return;
    }
    addNodeBetween(n->prev, n, d);
    l->size++;    
}

// search looks for Node on the pos. If pos doesn't from [0;size-1], then returns NULL
Node* search(List* l, int pos) {
    if (pos > l->size-1 || pos < 0) {
        return NULL;
    }

    Node* n = newNode(NULL); 
    if (pos >= l->size/2) {
        n = l->tail;        
        for (int i = 0; i < l->size-pos; i++) {
            n = n->prev;
        }     
    } else {
        n = l->head;                
        for (int i = 0; i <= pos; i++) {
            n = n->next;
        }
    }
    return n;
}

// erase deletes Node on the pos. If pos is invalid, then returns NULL
void erase(List* l, int pos) {
    Node* n = search(l, pos);
    if (n == NULL) {
        printf("search returns NULL\n");        
        return;
    }

    deleteNode(n->prev, n, n->next);
    l->size--;
}

// printList prints List with head, all elems and tail
void printList(List* l) {
    printf("HEAD <-");

    Node* n = l->head;    
    for (int i = 0; i < l->size; i++) {
        printf("-> ");
        print(n->next->d, Brief);
        printf(" <-");
        n = n->next;    
    }

    printf("-> TAIL\n");
}

// freeList deletes List and appropriate frees memory
void freeList(List* l) {
    Node* n = l->head;    
    for (int i = 0; i <= l->size; i++) {
        n = n->next;    
        free(n->prev);
    }
    free(n);
    free(l);
}

static void addNodeBetween(Node* lhs, Node* rhs, Date* d) {
    Node* n = newNode(d);
    n->next = rhs;
    n->prev = lhs;    
    
    lhs->next = n;
    rhs->prev = n;
}

static Node* newNode(Date* d) {
    Node* n = malloc(sizeof(Node));
    n->d = d;
    n->next = NULL;
    n->prev = NULL;

    return n;
}

static void deleteNode(Node* lhs, Node* n, Node* rhs) {
    lhs->next = rhs;
    rhs->prev = lhs;

    n->next = NULL;
    n->prev = NULL;

    free(n);
}
