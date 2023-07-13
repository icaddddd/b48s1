const janji = new Promise((resolve, reject) => {
    const xhttp = new XMLHttpRequest()

    xhttp.open("GET", "https://api.npoint.io/b5f129812bd0f453b24e", true)
    xhttp.onload = function () {
        if (xhttp.status === 200) {
            resolve(JSON.parse(xhttp.responseText))
        } else if (xhttp.status >= 400) {
            reject("Data Error!!!")
        }
    }
    xhttp.onerror = function() {
        reject("Network Error!!!")
    }
    xhttp.send()
})

let testimonialGroup = []

async function getGroup (rating) {
    try {
        const response = await janji
        console.log(response)
        testimonialGroup = response
        buttonTestimonial()
    } catch (err) {
        console.log(err)
    }
}

getGroup()

function buttonTestimonial() {  
    let testimonialAja = ""

    testimonialGroup.forEach((kotak) => {
        testimonialAja += `<div class="container-grid-lagi" id="container-grid-lagi">
        <img src="${kotak.image}">
        <p class="quote">${kotak.quote}</p>
        <p class="user">- ${kotak.user}</p>
        <p class="user">${kotak.rating} <i class="fa-sharp fa-regular fa-star" style="color: #ff4500;"></i> </p>
        </div>`
    })

    document.getElementById("container-grid").innerHTML = testimonialAja
}

function ratingTestimonial(rating) {
    let filterTestimonialAja = ""

    const udahdiFilter = testimonialGroup.filter((kotak) => {
        return kotak.rating === rating
    })

    udahdiFilter.forEach((kotak) => {
        filterTestimonialAja += `<div class="container-grid-lagi" id="container-grid-lagi">
        <img src="${kotak.image}">
        <p class="quote">${kotak.quote}</p>
        <p class="user">- ${kotak.user}</p>
        <p class="user">${kotak.rating} <i class="fa-sharp fa-regular fa-star" style="color: #ff4500;"></i> </p>
        </div>`
    })

    document.getElementById("container-grid").innerHTML = filterTestimonialAja
}   