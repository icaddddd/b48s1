function submitData(event){
    event.preventDefault()

    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phonenumber").value
    let select = document.getElementById("select").value
    let description = document.getElementById("description").value

    let objectData = {
        name,
        email,
        phone,
        select,
        description
    }

    if (name === ""){
        alert('nama harus diisi')
    } else if (email === ""){
        return alert('email harus diisi')
    } else if (phone === ""){
        return alert('phone harus diisi')
    } else if (select === ""){
        return alert('subject harus diisi')
    } else if (description === ""){
        return alert('message harus diisi')
    }

    const emailReceiver = "rhisjaddjvtr@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${select}&body=Halo nama saya ${name},\n${description}, silahkan kontak saya di nomor berikut : ${phone}`
    a.click()
}