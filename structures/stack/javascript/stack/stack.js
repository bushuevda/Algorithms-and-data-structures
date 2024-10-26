class Stack{
    constructor(){
        this.array_ = Array();
        this.type_ = "Stack";
    }

    /*
    A function that returns a stack array
    */
    get_array(){
        return this.array_;
    }

    /*
    A function that returns the top element of the stack 
    */
    top(){
        if (! this.empty())
            return this.array_[this.array_.length - 1];
        else
            console.log("size of the stack must be more than 0")
    }
    
    /*
    A function that checks the stack for emptiness
     */
    empty(){
        if (this.size() > 0)
            return false;
        else 
            return true;
    }

    /*
    A function that removes a top element from the stack and returns it
    */
    pop(){
        if (! this.empty()){
            let temp = this.array_[this.array_.length - 1];
            this.array_ = this.array_.slice(0, this.array_.length - 1);
            return temp;
        }
        else
            console.log("size of the stack must be more than 0")
    }

    /*
    A function that adds an element to the stack
    */
    push(element){
        this.array_.push(element);
    }

    /*
    A function that returns stack size
     */
    size(){
        return this.array_.length;
    }

    /*
    A function that swaps the stack arrays
     */
    swap(stack){
        if (stack.type_ == this.type_){
            let temp = this.array_;
            this.array_ = stack.array_;
            stack.array_ = temp;
        }
        else
            console.log("type of the stack must be 'Stack'")
    }

    /*
    A function that shows the stack arrays
    */
    show(){
        console.log(this.array_)
    }
}

export {Stack};