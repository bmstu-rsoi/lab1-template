namespace RSOI_Lab_01.Models
{
    public class ErrorResponse
    {
        /// <summary>
        /// Сообщение об ошибке
        /// </summary>
        public string? Message { get; set; }

        public ErrorResponse(string? message)
        {
            Message = message;
        }
    }
}
