async function apiRequest(url, options = {}) {

    try {

        const response = await fetch(url, options);

        const result = await response.json();

        if (!result.status) {
            throw new Error(result.message_error);
        }

        return result.data;

    } catch(error) {

        console.error(error);
        alert(error.message);

        return null;
    }
}



document
.getElementById("loadShop")
.onclick = async()=>{

    const shops = await apiRequest("/shops/viewshoplist");

    if(!shops) return;

    
    const container =
        document.getElementById("shopContainer");

    container.innerHTML="";


    shops.forEach(shop=>{

        const div=document.createElement("div");

        div.className="shop";

        div.innerHTML=`

            <h3>${shop.name}</h3>

            <p>
            Location:
            ${shop.location}
            </p>

            <p>
            ID:
            ${shop.id}
            </p>

            <button onclick="loadItems('${shop.id}')">
                View Items
            </button>

        `;

        container.appendChild(div);

    });

}


async function loadItems(shopId){

    const shop = await apiRequest(
        `/shops/${shopId}/viewitem`
    );


    if(!shop) return;


    const container =
        document.getElementById("shopContainer");


    container.innerHTML = `

        <h3>
            ${shop.name}
        </h3>

        <p>
            Location:
            ${shop.location}
        </p>

    `;



    shop.items.forEach(item=>{


        const div =
        document.createElement("div");


        div.className="item";


        div.innerHTML=`
            <h4>
                ${item.name}
            </h4>


            <p>
                Stock:
                ${item.stock}
            </p>


            <label>
                Availability:
            </label>


            <select 
                onchange="changeAvailability('${item.id}', this.value)"
            >


            <option 
            value="available"
            ${item.available === "available" ? "selected" : ""}
            >
            Available
            </option>


            <option 
            value="not_available"
            ${item.available === "not_available" ? "selected" : ""}
            >
            Not Available
            </option>


            </select>


            <br>


            <button onclick="
                addStock('${item.id}')
            ">
                Add Stock
            </button>


            <button onclick="
                checkout('${item.id}')
            ">
                Checkout
            </button>



        `;


        container.appendChild(div);


    });

}



document
.getElementById("addProduct")
.onsubmit = async(e)=>{

    e.preventDefault();


    const form = new FormData();


    form.append(
        "name",
        document.getElementById("name").value
    );


    form.append(
        "price",
        document.getElementById("price").value
    );


    form.append(
        "stock",
        document.getElementById("stock").value
    );


    form.append(
        "available",
        document.getElementById("available").value
    );


    form.append(
        "shop_id",
        document.getElementById("shopId").value
    );



    const result = await apiRequest(
        "/products/additem",
        {
            method:"POST",
            body:form
        }
    );


    if(result){

        alert("Product added successfully");

    }

};

async function addStock(itemId){


    const amount =
        prompt(
            "Add stock amount:"
        );


    if(!amount)
        return;



    const form = new FormData();


    form.append(
        "item_id",
        itemId
    );


    form.append(
        "stock",
        amount
    );



    const result =
    await apiRequest(
        "/products/additemstock",
        {
            method:"PUT",
            body:form
        }
    );



    if(result !== null){

        alert(
            "Stock added"
        );

        document
        .getElementById("loadShop")
        .click();

    }

}

async function checkout(itemId){


    const quantity =
        prompt(
            "Checkout quantity:"
        );


    if(!quantity)
        return;



    const body = {

        quantity:
        Number(quantity)

    };



    const result =
    await apiRequest(
        `/products/${itemId}/checkout`,
        {

            method:"POST",

            headers:{
                "Content-Type":
                "application/json"
            },


            body:
            JSON.stringify(body)

        }
    );



    if(result !== null){

        alert(
            "Checkout success"
        );


        document
        .getElementById("loadShop")
        .click();

    }

}

async function changeAvailability(itemId, currentStatus){


    let newStatus;


    if(currentStatus === "available"){

        newStatus = "not_available";

    }else{

        newStatus = "available";

    }



    const form = new FormData();


    form.append(
        "id",
        itemId
    );


    form.append(
        "availabilty",
        newStatus
    );



    const result =
    await apiRequest(
        "/products/updateavailabilty",
        {
            method:"PUT",
            body:form
        }
    );



    if(result !== null){

        alert(
            `Item changed to ${newStatus}`
        );


        // refresh data
        document
        .getElementById("loadShop")
        .click();

    }


}

document
.getElementById("addShop")
.onsubmit = async(e)=>{


    e.preventDefault();


    const form = new FormData();


    form.append(
        "name",
        document
        .getElementById("shopName")
        .value
    );


    form.append(
        "location",
        document
        .getElementById("shopLocation")
        .value
    );



    const result =
    await apiRequest(
        "/shops/addshop",
        {
            method:"POST",
            body:form
        }
    );



    if(result !== null){

        alert(
            "Shop added successfully"
        );


        // clear form

        document
        .getElementById("shopName")
        .value="";


        document
        .getElementById("shopLocation")
        .value="";


        // refresh shop list

        document
        .getElementById("loadShop")
        .click();

    }

};