(require 'elisp-mode)
(require 'generic)

(define-derived-mode lavaloon-mode emacs-lisp-mode "Lavaloon"
  "Major-mode for editing lavaloon lisp code.")

(add-to-list 'auto-mode-alist '("\\.lv\\'"  . lavaloon-mode))

(provide 'lavaloon-mode)
