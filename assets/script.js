new WOW().init();
const swiperEl = document.querySelector('.mySwiper')
const params = {
    autoplay: {
        delay: 2500,
    },
    loop: false,
    slidesPerView: 1,
    spaceBetween: 0,
    enteredSlides: true,
    roundLengths: true,
    grabCursor: true,
    injectStyles: [`
          .swiper-pagination-bullet {
            width: 10px;
            height: 10px;
            background: #1A1A1A;
            border: 1px solid #1A1A1A;
            border-radius: 50%;
            opacity: 1;
            box-sizing: border-box;
          }
    
          .swiper-pagination-bullet-active {
            background: #fff;
            border: 1px solid #1A1A1A;
          }
          `],
    pagination: {
        el: '.swiper-pagination',
        clickable: true,
        renderBullet: function (index, className) {
            return '<span class="' + className + '">' + "</span>";
        },
    },
    navigation: {
        nextEl: '.swiper-button-next',
        prevEl: '.swiper-button-prev',
    },
    breakpoints: {
        768: {
            slidesPerView: 1.5,
            spaceBetween: 16,
        },

    },
}

Object.assign(swiperEl, params)
swiperEl.initialize();