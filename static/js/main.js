document.addEventListener("DOMContentLoaded", () => {
  // ==========================================================================
  // 1. GESTION DE L'OUVERTURE / FERMETURE DE LA MODAL
  // ==========================================================================
  const modal = document.getElementById("createPostModal");
  const openModalBtn = document.querySelector(".create-input-link");
  const closeModalBtn = document.getElementById("closeModalBtn");
  const cancelModalBtn = document.getElementById("cancelModalBtn");

  // Ouvre la modal au clic sur la barre
  if (openModalBtn && modal) {
    openModalBtn.addEventListener("click", (e) => {
      e.preventDefault();
      modal.classList.add("active");
      document.body.style.overflow = "hidden";
    });
  }

  const closeModal = () => {
    if (modal) {
      modal.classList.remove("active");
      document.body.style.overflow = "";
    }
  };

  if (closeModalBtn) closeModalBtn.addEventListener("click", closeModal);
  if (cancelModalBtn) cancelModalBtn.addEventListener("click", closeModal);

  if (modal) {
    modal.addEventListener("click", (e) => {
      if (e.target === modal) {
        closeModal();
      }
    });
  }

  // ==========================================================================
  // 2. GESTION DYNAMIQUE DES BADGES DE CATÉGORIES (DATALIST)
  // ==========================================================================
  const searchInput = document.getElementById("categorySearchInput");
  const datalist = document.getElementById("categoriesData");
  const badgesContainer = document.getElementById("selectedCategoriesBadges");
  const form = document.getElementById("createPostForm");

  const selectedCategoryIds = new Set();

  if (searchInput && datalist && badgesContainer) {
    searchInput.addEventListener("input", function () {
      const value = this.value;
      const options = datalist.querySelectorAll("option");

      options.forEach((option) => {
        if (option.value === value) {
          const catId = option.getAttribute("data-id");
          const catName = option.value;

          if (!selectedCategoryIds.has(catId)) {
            selectedCategoryIds.add(catId);

            const badge = document.createElement("span");
            badge.className = "category-badge";
            badge.innerHTML = `${catName} <span class="remove-badge" data-id="${catId}">&times;</span>`;

            const hiddenInput = document.createElement("input");
            hiddenInput.type = "hidden";
            hiddenInput.name = "categories";
            hiddenInput.value = catId;
            hiddenInput.id = `hidden-cat-${catId}`;

            badge.appendChild(hiddenInput);
            badgesContainer.appendChild(badge);
          }

          this.value = "";
        }
      });
    });

    // Écouteur pour supprimer un badge quand on clique sur sa petite croix
    badgesContainer.addEventListener("click", function (e) {
      if (e.target.classList.contains("remove-badge")) {
        const catId = e.target.getAttribute("data-id");
        selectedCategoryIds.delete(catId);
        e.target.parentElement.remove();
      }
    });
  }

  // ==========================================================================
  // 3. SÉCURITÉ SUBMIT : REQUISITION D'AU MOINS UNE CATÉGORIE
  // ==========================================================================
  if (form) {
    form.addEventListener("submit", function (e) {
      if (selectedCategoryIds.size === 0) {
        e.preventDefault();
        alert("Veuillez ajouter au moins une catégorie à votre publication !");
      }
    });
  }
});
