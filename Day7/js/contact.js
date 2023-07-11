function submitData(event) {
    event.preventDefault()

    let name = document.getElementById("inputname").value
    let email = document.getElementById("inputemail").value
    let phone = document.getElementById("inputphone").value
    let subject = document.getElementById("inputsubject").value
    let message = document.getElementById("inputmessage").value


    let objectData = {
        name,
        email,
        phone,
        subject,
        message
    }

    console.log(objectData)

    if (name === ""){
        alert('nama harus diisi')
    } else if (email === ""){
        return alert('email harus diisi')
    } else if (phone === ""){
        return alert('phone harus diisi')
    } else if (subject === ""){
        return alert('subject harus diisi')
    } else if (message === ""){
        return alert('message harus diisi')
    }

    const emailReceiver = "rhisjaddjvtr@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${phone}`
    a.click()

}