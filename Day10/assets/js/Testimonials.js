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
        testimonialAja += `
        
              <div class="col-auto d-flex align-items-center mb-5" id="card450">
                <div class="card" style="width: 18rem; height: 400px;">
                    <img src="${kotak.image}" class="card-img-top object-fit-cover" style="height: 230px;" id="image">
                    <div class="card-body">
                      <p class="card-text text-start overflow-hidden white-space-wrap" id="quote" style="height: 50px; font-size: 12px;font-style: italic;">${kotak.quote}</p>
                      <p class="card-text text-end" id="author" style="font-size: 16px;">- ${kotak.user}</p>
                      <p class="card-text text-end" id="rating">${kotak.rating} <i class="fa-sharp fa-regular fa-star" style="color: #ff4500;"></i></p>
                    </div>
                </div>
              </div>
        
    `
    })

    document.getElementById("card-testimonials").innerHTML = testimonialAja
}

function ratingTestimonial(rating) {
    let filterTestimonialAja = ""

    const udahdiFilter = testimonialGroup.filter((kotak) => {
        return kotak.rating === rating
    })

    udahdiFilter.forEach((kotak) => {
        filterTestimonialAja += `
        
              <div class="col-auto d-flex align-items-center mb-5" id="card450">
                <div class="card" style="width: 18rem; height: 400px;">
                    <img src="${kotak.image}" class="card-img-top object-fit-cover" style="height: 230px;" id="image">
                    <div class="card-body">
                      <p class="card-text text-start overflow-hidden white-space-wrap" id="quote" style="height: 50px; font-size: 12px;font-style: italic;">${kotak.quote}</p>
                      <p class="card-text text-end" id="author" style="font-size: 16px;">- ${kotak.user}</p>
                      <p class="card-text text-end" id="rating">${kotak.rating} <i class="fa-sharp fa-regular fa-star" style="color: #ff4500;"></i></p>
                    </div>
                </div>
              </div>
       
    `
    })

    document.getElementById("card-testimonials").innerHTML = filterTestimonialAja
}