// Sidebar functionality
document.addEventListener('DOMContentLoaded', function() {
    // Get the sidebar element
    const sidebar = document.getElementById('logo-sidebar');
    
    // Function to hide the sidebar (for small screens)
    function hideSidebar() {
        if (sidebar && window.innerWidth < 1024) {
            sidebar.classList.remove('translate-x-0');
            sidebar.classList.add('-translate-x-full');
        }
    }
    
    // Function to show the sidebar (for small screens)
    function showSidebar() {
        if (sidebar && window.innerWidth < 1024) {
            sidebar.classList.add('translate-x-0');
            sidebar.classList.remove('-translate-x-full');
        }
    }
    
    // Function to toggle the sidebar
    function toggleSidebar(event) {
        // Prevent default and stop propagation to avoid interference
        if (event) {
            event.preventDefault();
            event.stopPropagation();
        }
        
        if (sidebar) {
            if (sidebar.classList.contains('-translate-x-full')) {
                showSidebar();
            } else {
                hideSidebar();
            }
        }
    }
    
    // Set up all click handlers
    function setupSidebarHandlers() {
        // 1. Toggle button in header
        const toggleButton = document.getElementById('sidebar-toggle');
        if (toggleButton) {
            // Remove any existing listeners first
            toggleButton.removeEventListener('click', toggleSidebar);
            // Add new listener
            toggleButton.addEventListener('click', toggleSidebar);
        }
        
        // 2. Elements with sidebar-close class (navigation links)
        const sidebarCloseElements = document.querySelectorAll('.sidebar-close');
        sidebarCloseElements.forEach(element => {
            // Remove any existing listeners first
            element.removeEventListener('click', hideSidebar);
            // Add new listener with capture to ensure it runs before other handlers
            element.addEventListener('click', function(event) {
                if (window.innerWidth < 1024) {
                    // Don't prevent default here as we want HTMX to work
                    hideSidebar();
                }
            }, true);
        });
        
        // 3. Special handling for elements with data-sidebar-toggle attribute (Alex link)
        const specialToggleElements = document.querySelectorAll('[data-sidebar-toggle]');
        specialToggleElements.forEach(element => {
            // Remove any existing listeners first
            element.removeEventListener('click', hideSidebar);
            // Add new listener with capture phase to ensure it runs first
            element.addEventListener('click', function(event) {
                if (window.innerWidth < 1024) {
                    // Don't prevent default, just hide the sidebar
                    // We specifically want HTMX to work
                    hideSidebar();
                }
            }, true);
        });
    }

    // Handle window resize to ensure correct sidebar state
    function handleResize() {
        if (window.innerWidth >= 1024) {
            // Ensure sidebar is visible on large screens
            if (sidebar) {
                sidebar.classList.remove('-translate-x-full');
                sidebar.classList.add('translate-x-0');
            }
        } else {
            // Ensure sidebar is hidden on small screens
            if (sidebar) {
                sidebar.classList.remove('translate-x-0');
                sidebar.classList.add('-translate-x-full');
            }
        }
    }

    // Initialize sidebar handlers
    setupSidebarHandlers();
    
    // Set up resize handler
    window.addEventListener('resize', handleResize);
    
    // Run handler once on load to ensure correct initial state
    handleResize();
    
    // Listen for HTMX events to reinitialize handlers
    document.body.addEventListener('htmx:afterSwap', function(event) {
        // Small delay to ensure DOM is fully updated
        setTimeout(setupSidebarHandlers, 100);
    });
    
    // Handle click outside sidebar to close it on mobile
    document.addEventListener('click', function(event) {
        // Only on mobile screens
        if (window.innerWidth < 1024 && sidebar) {
            // If sidebar is visible and click is outside sidebar
            if (!sidebar.classList.contains('-translate-x-full') && 
                !sidebar.contains(event.target) && 
                !event.target.closest('#sidebar-toggle')) {
                hideSidebar();
            }
        }
    });
});