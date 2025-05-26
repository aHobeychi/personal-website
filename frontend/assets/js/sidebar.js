/**
 * Alpine.js sidebar integration with HTMX
 * Maintains sidebar state during HTMX navigation
 */
document.addEventListener('alpine:init', () => {
    // Global sidebar store
    Alpine.store('sidebar', {
        open: false,
        
        toggle() {
            this.open = !this.open;
        },
        
        close() {
            this.open = false;
        }
    });
});

// Handle HTMX navigation - ensure Alpine components are re-initialized
document.body.addEventListener('htmx:afterSwap', function(event) {
    // Close sidebar on navigation for mobile
    if (window.innerWidth < 1024 && Alpine.store('sidebar')) {
        Alpine.store('sidebar').close();
    }
});

// Handle window resize - close sidebar when switching to desktop
window.addEventListener('resize', function() {
    if (window.innerWidth >= 1024 && Alpine.store('sidebar')) {
        Alpine.store('sidebar').close();
    }
});

// Close sidebar when clicking outside on mobile
document.addEventListener('click', function(event) {
    if (window.innerWidth < 1024 && Alpine.store('sidebar') && Alpine.store('sidebar').open) {
        const sidebar = document.querySelector('aside[role="complementary"]');
        const toggleButton = document.querySelector('button[aria-label="Toggle sidebar menu"]');
        
        if (sidebar && !sidebar.contains(event.target) && 
            toggleButton && !toggleButton.contains(event.target)) {
            Alpine.store('sidebar').close();
        }
    }
});