document.addEventListener("DOMContentLoaded", function () {
  const submitBtn = document.getElementById("submit-btn");

  submitBtn.addEventListener("click", function (e) {
    e.preventDefault();

    const inputs = document.querySelectorAll(
      '.input-fields input[type="text"]',
    );
    let allFilled = true;

    inputs.forEach((input) => {
      if (!input.value.trim()) {
        allFilled = false;
        input.style.border = "2px solid red";
      } else {
        input.style.border = "";
      }
    });

    if (allFilled) {
      window.location.href = "../acceptPage.html";
    }
  });
});
