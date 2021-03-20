namespace Genocs.Microservice.Models
{
    public class OrderPayload
    {
        public Order Data { get; set; }
    }

    public class Order
    {
        public string Id { get; set; }
        public string Description { get; set; }
    }

    public class OrderConfirmationPayload
    {
        public OrderConfirmation Data { get; set; }
    }

    public class OrderConfirmation
    {
        public string Id { get; set; }
        public string Description { get; set; }
        public string Status { get; set; }
    }
}
