class Stack<T>{
    protected array_: Array<T> = new Array<T>();
    
    /*
    A function that returns a stack array
     */
    getArray(): T[]{
        return this.array_;
    }

    /*
    A function that returns stack size
     */
    size(): number{
        return this.array_.length;
    }

    /*
    A function that checks the stack for emptiness
     */
    empty(): boolean{
        return this.size() > 0 ? false : true;
    }

    /*
    A function that adds an element to the stack
    */
    push(element: T){
        this.array_.push(element);
    }

    /*
    A function that returns the top element of the stack 
    */
    top(): T | undefined{
        let temp: T | undefined = undefined;
        if(!this.empty()){
            temp = this.array_[this.array_.length - 1];
        }
        return temp;
    }

    /*
    A function that removes a top element from the stack and returns it
    */
    pop(): T | undefined{
        let temp: T | undefined = undefined;
        if(!this.empty()){
            temp = this.array_[this.array_.length - 1];
            this.array_ = this.array_.slice(0, this.array_.length - 1);
        }
        return temp;
    }

    /*
    A function that swaps the stack arrays
     */
    swap(s2: Stack<T>){
        let temp: Array<T> = new Array<T>();
        temp = s2.array_;
        s2.array_ = this.array_;
        this.array_ = temp;
    }

    /*
    A function that shows the stack arrays
    */
    show(){
        console.log(this.array_);
    }

}

export {Stack};