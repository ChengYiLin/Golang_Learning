console.log("Hello JavaScript")

const signUpBtn = document.getElementById("signUpBtn")
const signInBtn = document.getElementById("signInBtn")
const signOut = document.getElementById("signOut")

// Sign Up
signUpBtn.addEventListener("click", (e) => {
    e.preventDefault()

    const name = document.getElementById('register_name').value
    const account = document.getElementById('register_account').value
    const passWord = document.getElementById('register_password').value

    const request = {
        method: 'POST',
        body:JSON.stringify({
          name: name,
          account: account,
          passWord: passWord
        }),
        headers: new Headers({
            'Content-Type': 'application/json'
        })
    }

    fetch("/signup", request)
        .then(res => res.json())
        .then(res => {
            if (res.Status === "error") {
                window.location.href = `/error?errorMsg=${res.ErrorMessage}`
            } else {
                window.location.href = `/member?name=${res.Name}`
            }
        })
        .catch(err => {
            alert("[POST Error] : /signup")
            console.log(err)
        })
})

// Sign In
signInBtn.addEventListener("click", (e) => {
    e.preventDefault()

    const account = document.getElementById('login_account').value
    const passWord = document.getElementById('login_password').value

    const request = {
        method: 'POST',
        body:JSON.stringify({
          account: account,
          passWord: passWord
        }),
        headers: new Headers({
            'Content-Type': 'application/json'
        })
    }

    fetch("/signin", request)
        .then(res => res.json())
        .then(res => {
            if (res.Status === "error") {
                window.location.href = `/error?errorMsg=${res.ErrorMessage}`
            } else {
                window.location.href = `/member?name=${res.Name}`
            }
        })
        .catch(err => {
            alert("[POST Error] : /signup")
            console.log(err)
        })
})

// Sign out
signOut.addEventListener("click", (e) => {
    e.preventDefault()

    window.location.href ="/"
})