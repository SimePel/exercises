#include "stdio.h"
#include "list.h"

int main() {
    List* l = newList();
    Date zero = {1, 1, 2017};    
    push_back(l, &zero);
    Date one = {2, 2, 2017};
    push_back(l, &one);
    Date two = {3, 3, 2017};    
    push_front(l, &two);
    Date three = {4, 4, 2017};        
    push_front(l, &three);

    printList(l);
      
    Date four = {5, 5, 2017};
    insert_after(l, 1, &four);
    Date five = {6, 6, 2017};
    insert_before(l, 4, &five);

    printList(l);
    
    erase(l,0);
    erase(l,4);
    erase(l,2);
    erase(l,2);
    erase(l,1);
    erase(l,0);

    printList(l);

    freeList(l);
    return 0;
}