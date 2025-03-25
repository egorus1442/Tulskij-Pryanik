document.addEventListener("DOMContentLoaded", function () {
  const cancelBtn = document.getElementById("cancel-btn");
  if (cancelBtn) {
    cancelBtn.addEventListener("click", function () {
      window.location.href = "../mainPage.html";
    });
  }

  const continueBtn = document.getElementById("continue-btn");
  if (continueBtn) {
    continueBtn.addEventListener("click", function () {
      window.location.href = "../divisionResultPage.html";
    });
  }
});
