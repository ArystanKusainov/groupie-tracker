const rangeInput1 = document.querySelectorAll(".range-input1 input"),
priceInput1 = document.querySelectorAll(".price-input1 input"),
range1 = document.querySelector(".slider1 .progress1");
let priceGap1 = 1;

priceInput1.forEach(input =>{
    input.addEventListener("input", e =>{
        let minPrice = parseInt(priceInput1[0].value),
        maxPrice = parseInt(priceInput1[1].value);
        
        if((maxPrice - minPrice >= priceGap1) && maxPrice <= rangeInput1[1].max){
            if(e.target.className === "input-min"){
                rangeInput1[0].value = minPrice;
                range1.style.left = ((minPrice / (4*(rangeInput1[0].max-75))) * 100) + "%";
            }else{
                rangeInput1[1].value = maxPrice;

                range1.style.right = 100 - 0.75*(maxPrice / (rangeInput1[1].max-25)) * 100 + "%";
            }
        }
    });
});

rangeInput1.forEach(input =>{
    input.addEventListener("input", e =>{
        let minVal = parseInt(rangeInput1[0].value),
        maxVal = parseInt(rangeInput1[1].value);

        if((maxVal - minVal) < priceGap1){
            if(e.target.className === "range-min"){
                rangeInput1[0].value = maxVal - priceGap1
            }else{
                rangeInput1[1].value = minVal + priceGap1;
            }
        }else{
            priceInput1[0].value = minVal;
            priceInput1[1].value = maxVal;
            range1.style.left = 100 -(((rangeInput1[0].max)-minVal) ) + "%";
            range1.style.right =     ((rangeInput1[1].max)-maxVal) + "%";
        }
    });
});


const rangeInput2 = document.querySelectorAll(".range-input2 input"),
priceInput2 = document.querySelectorAll(".price-input2 input"),
range2 = document.querySelector(".slider2 .progress2");
let priceGap2 = 1;

priceInput2.forEach(input =>{
    input.addEventListener("input", e =>{
        let minPrice = parseInt(priceInput2[0].value),
        maxPrice = parseInt(priceInput2[1].value);
        
        if((maxPrice - minPrice >= priceGap2) && maxPrice <= rangeInput2[1].max){
            if(e.target.className === "input-min"){
                rangeInput2[0].value = minPrice;
                range2.style.left = ((minPrice / (4*(rangeInput2[0].max-75))) * 100) + "%";
            }else{
                rangeInput2[1].value = maxPrice;

                range2.style.right = 100 - 0.75*(maxPrice / (rangeInput2[1].max-25)) * 100 + "%";
            }
        }
    });
});

rangeInput2.forEach(input =>{
    input.addEventListener("input", e =>{
        let minVal = parseInt(rangeInput2[0].value),
        maxVal = parseInt(rangeInput2[1].value);

        if((maxVal - minVal) < priceGap2){
            if(e.target.className === "range-min"){
                rangeInput2[0].value = maxVal - priceGap2
            }else{
                rangeInput2[1].value = minVal + priceGap2;
            }
        }else{
            priceInput2[0].value = minVal;
            priceInput2[1].value = maxVal;
            range2.style.left = 100 -(((rangeInput2[0].max)-minVal) ) + "%";
            range2.style.right =     ((rangeInput2[1].max)-maxVal) + "%";
        }
    });
});
