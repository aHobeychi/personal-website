/**
 * Sidebar functionality for responsive navigation
 * Uses a self-executing function to create a closure
 * and prevent global namespace pollution
 */
(function() {
    // Create a namespace for our sidebar functionality
    window.SidebarManager = window.SidebarManager || {
        isInitialized: false,
        eventHandlers: {},
        init: function() {
            // Prevent multiple initializations
            this.cleanup();
            this.isInitialized = true;
            this.setupEventListeners();
            this.handleInitialState();
        },
        cleanup: function() {
            // Clean up all registered event listeners
            if (this.eventHandlers.resize) {
                window.removeEventListener('resize', this.eventHandlers.resize);
            }
            if (this.eventHandlers.outsideClick) {
                document.removeEventListener('click', this.eventHandlers.outsideClick);
            }
            
            // Clean up toggle button listeners
            const toggleButton = document.getElementById('sidebar-toggle');
            if (toggleButton && this.eventHandlers.toggle) {
                toggleButton.removeEventListener('click', this.eventHandlers.toggle);
            }
            
            // Clean up sidebar close element listeners
            const sidebarCloseElements = document.querySelectorAll('.sidebar-close');
            if (sidebarCloseElements && this.eventHandlers.close) {
                sidebarCloseElements.forEach(element => {
                    element.removeEventListener('click', this.eventHandlers.close);
                });
            }
            
            // Reset the handlers object
            this.eventHandlers = {};
            this.isInitialized = false;
        },
        setupEventListeners: function() {
            const sidebar = document.getElementById('logo-sidebar');
            if (!sidebar) {
                console.error('Sidebar element not found');
                return;
            }
            
            // Define hide, show, and toggle functions
            this.eventHandlers.hide = () => {
                if (sidebar && window.innerWidth < 1024) {
                    sidebar.classList.remove('translate-x-0');
                    sidebar.classList.add('-translate-x-full');
                }
            };
            
            this.eventHandlers.show = () => {
                if (sidebar && window.innerWidth < 1024) {
                    sidebar.classList.add('translate-x-0');
                    sidebar.classList.remove('-translate-x-full');
                }
            };
            
            this.eventHandlers.toggle = (event) => {
                if (event) {
                    event.preventDefault();
                    event.stopPropagation();
                }
                
                if (sidebar) {
                    if (sidebar.classList.contains('-translate-x-full')) {
                        this.eventHandlers.show();
                    } else {
                        this.eventHandlers.hide();
                    }
                }
            };
            
            // Set up toggle button handler
            const toggleButton = document.getElementById('sidebar-toggle');
            if (toggleButton) {
                toggleButton.addEventListener('click', this.eventHandlers.toggle);
            }
            
            // Set up close handlers for navigation links
            this.eventHandlers.close = (event) => {
                if (window.innerWidth < 1024) {
                    this.eventHandlers.hide();
                }
            };
            
            const sidebarCloseElements = document.querySelectorAll('.sidebar-close');
            sidebarCloseElements.forEach(element => {
                element.addEventListener('click', this.eventHandlers.close, true);
            });
            
            // Handle outside click to close sidebar
            this.eventHandlers.outsideClick = (event) => {
                if (window.innerWidth < 1024 && sidebar && 
                    !sidebar.classList.contains('-translate-x-full') && 
                    !sidebar.contains(event.target) && 
                    !event.target.closest('#sidebar-toggle')) {
                    this.eventHandlers.hide();
                }
            };
            document.addEventListener('click', this.eventHandlers.outsideClick);
            
            // Resize handler
            this.eventHandlers.resize = () => {
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
            };
            window.addEventListener('resize', this.eventHandlers.resize, { passive: true });
        },
        handleInitialState: function() {
            // Set initial state based on screen size
            if (this.eventHandlers.resize) {
                this.eventHandlers.resize();
            }
        }
    };
    
    // Initialize the sidebar manager
    function initializeSidebar() {
        // Use setTimeout to ensure DOM is fully loaded
        setTimeout(function() {
            // Only initialize if necessary elements exist
            if (document.getElementById('logo-sidebar')) {
                window.SidebarManager.init();
            }
        }, 50);
    }
    
    // Set up initialization on various events
    
    // On DOM content loaded
    document.addEventListener('DOMContentLoaded', initializeSidebar);
    
    // On HTMX after swap (page navigation)
    document.body.addEventListener('htmx:afterSwap', initializeSidebar);
    
    // On HTMX after settle (all animations completed)
    document.body.addEventListener('htmx:afterSettle', initializeSidebar);
    
    // On browser navigation (back/forward)
    window.addEventListener('popstate', initializeSidebar);
    
    // Expose reinit function for debugging
    window.reinitSidebar = initializeSidebar;
    
    // Initial call (for when script loads after DOM is already ready)
    if (document.readyState === 'complete' || document.readyState === 'interactive') {
        initializeSidebar();
    }
})();