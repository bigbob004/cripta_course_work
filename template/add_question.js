let question_cnt = 2
function add_question(){
    let block = document.querySelector(".form_group")
    let clone = block.cloneNode(true)
    clone.querySelector('.number').innerText = question_cnt
    block.parentElement.appendChild(clone)
    ++question_cnt
}