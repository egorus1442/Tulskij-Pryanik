document.addEventListener("DOMContentLoaded", function () {
  const calculateBtn = document.getElementById("calculate-btn");

  if (calculateBtn) {
    calculateBtn.addEventListener("click", function () {
      document.getElementById("form-section").scrollIntoView({
        behavior: "smooth",
      });
    });
  }
});
